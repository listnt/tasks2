package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/listnt/tasks2/T2.2.11/internal/models"
)

const (
	layoutISO = "2006-01-02"
)

func (rep *repository) loadToCache() error {
	query := `select events.user_id,events.date,events.event,events.description from events `
	rows, err := rep.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		p := models.Event{}
		if err := rows.Scan(&p.UserId, &p.Date, &p.Event, &p.Description); err != nil {
			rep.lg.Error(err)
			continue
		}
		if rep.cache[p.UserId] == nil {
			rep.cache[p.UserId] = make(map[string][]models.Event)
		}
		rep.cache[p.UserId][p.Date] = append(rep.cache[p.UserId][p.Date], p)
	}
	rep.lg.Infof("Loaded cache, len:%d", len(rep.cache))
	return err
}
func (rep *repository) GetEvents(user_id int, date string) ([]models.Event, error) {
	dist := strings.Split(date, "Till")
	if len(dist) == 1 {
		return rep.cache[user_id][dist[0]], nil
	} else {
		var res []models.Event
		T1, err := time.Parse(layoutISO, dist[0])
		if err != nil {
			rep.lg.Error(err)
			return nil, err
		}
		T2, _ := time.Parse(layoutISO, dist[1])
		if err != nil {
			rep.lg.Error(err)
			return nil, err
		}
		fmt.Println(T1, T2)
		for t := T1; t.Before(T2) || T2.Equal(t); t = t.AddDate(0, 0, 1) {
			fmt.Println(t)
			res = append(res, rep.cache[user_id][t.Format(layoutISO)]...)
		}
		return res, nil
	}

}

func (rep *repository) Store(v models.Event) error {
	ord := v
	query := ""
	query = fmt.Sprintf("call insert_into_events(%d, '%s', '%s','%s')",
		ord.UserId, ord.Date, ord.Event, ord.Description)
	ExecRes, err := rep.db.Exec(query)
	if err != nil {
		rep.lg.Error(err)
		return err
	}
	rep.lg.Println(query, ExecRes, err)

	var myord models.Event
	myord.UserId = ord.UserId
	myord.Date = ord.Date
	myord.Event = ord.Event
	myord.Description = ord.Description
	if rep.cache[myord.UserId] == nil {
		rep.cache[myord.UserId] = make(map[string][]models.Event)
	}
	rep.cache[ord.UserId][ord.Date] = append(rep.cache[ord.UserId][ord.Date], myord)
	rep.lg.Infof("Added to cache: %s", ord.Event)
	return nil
}

func (rep *repository) UpdateEvent(user_id int, date string, event string, newValue models.Event) error {
	query := fmt.Sprintf("with cte as (select * from events where user_id=%d and date='%s' and event='%s' limit 1) update events s set date='%s',event='%s',description='%s' from cte where cte.id=s.id",
		user_id, date, event, newValue.Date, newValue.Event, newValue.Description)
	ExecRes, err := rep.db.Exec(query)
	if err != nil {
		rep.lg.Error(err)
		return err
	}
	for i, e := range rep.cache[user_id][date] {
		if e.Event == event {
			rep.cache[user_id][date][i] = rep.cache[user_id][date][len(rep.cache[user_id][date])-1]
			rep.cache[user_id][date] = rep.cache[user_id][date][:len(rep.cache[user_id][date])]
			break
		}
	}
	rep.cache[user_id][newValue.Date] = append(rep.cache[user_id][newValue.Date], newValue)

	rep.lg.Println(query, ExecRes, err)
	return nil
}

func (rep *repository) Delete(user_id int, date string, event string) error {
	query := fmt.Sprintf("delete from events where ctid in (select ctid from events where user_id=%d and date='%s' and event='%s' limit 1)  ",
		user_id, date, event)
	ExecRes, err := rep.db.Exec(query)
	if err != nil {
		rep.lg.Error(err)
		return err
	}
	for i, e := range rep.cache[user_id][date] {
		if e.Event == event {
			rep.cache[user_id][date][i] = rep.cache[user_id][date][len(rep.cache[user_id][date])-1]
			rep.cache[user_id][date] = rep.cache[user_id][date][:len(rep.cache[user_id][date])]
			break
		}
	}
	rep.lg.Println(query, ExecRes, err)
	return nil
}

func (rep *repository) Close() {
	rep.db.Close()
	rep.lg.Info("[!!!] DB closed")
}
