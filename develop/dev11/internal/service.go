package internal

import (
	"time"

	Config "github.com/listnt/tasks2/T2.2.11/config"
	"github.com/listnt/tasks2/T2.2.11/internal/models"
	"github.com/listnt/tasks2/T2.2.11/internal/repository"
	"github.com/sirupsen/logrus"
)

// Service ...
type Service interface {
	CreateEvent(event models.Event) error
	DeleteEvent(user_id int, date string, event string) error
	UpdateEvent(user_id int, date string, event string, newEvent models.Event) error
	EventForDay(user_id int, date string) ([]models.Event, error)
	EventsForWeek(user_id int, date string) ([]models.Event, error)
	EventForMonth(user_id int, date string) ([]models.Event, error)
	Close()
}

type service struct {
	rep repository.Repository
	lg  *logrus.Logger
}

// NewService ...
/*Инициализирует связь с БД.
Загружает данные из БД в кэш.
Подписывается на nats
*/
func NewService(lg *logrus.Logger, cfg *Config.Config) Service {
	var err error
	res := new(service)
	res.lg = lg
	res.rep, err = repository.NewRepository(lg, cfg)
	if err != nil {
		res.lg.Fatal(err)
	}
	return res
}

const (
	layoutISO = "2006-01-02"
)

func (svc *service) Close() {
	svc.rep.Close()
}

func (svc *service) CreateEvent(event models.Event) error {
	return svc.rep.Store(event)
}
func (svc *service) DeleteEvent(userId int, date string, event string) error {
	return svc.rep.Delete(userId, date, event)
}
func (svc *service) UpdateEvent(userId int, date string, event string, newEvent models.Event) error {
	return svc.rep.UpdateEvent(userId, date, event, newEvent)
}
func (svc *service) EventForDay(userId int, date string) ([]models.Event, error) {
	return svc.rep.GetEvents(userId, date)
}
func (svc *service) EventsForWeek(userId int, date string) ([]models.Event, error) {
	t, err := time.Parse(layoutISO, date)
	if err != nil {
		svc.lg.Error(err)
	}
	T1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	T1 = T1.AddDate(0, 0, -int(T1.Weekday()))
	T2 := T1
	T2 = T2.AddDate(0, 0, 7)
	return svc.rep.GetEvents(userId, T1.Format(layoutISO)+"Till"+T2.Format(layoutISO))
}
func (svc *service) EventForMonth(userId int, date string) ([]models.Event, error) {
	t, err := time.Parse(layoutISO, date)
	if err != nil {
		svc.lg.Error(err)
		return nil, err
	}
	T1 := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	T2 := T1
	T2 = T2.AddDate(0, 1, -1)
	return svc.rep.GetEvents(userId, T1.Format(layoutISO)+"Till"+T2.Format(layoutISO))
}
