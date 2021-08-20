package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //Postgres
	Config "github.com/listnt/tasks2/T2.2.11/config"
	"github.com/listnt/tasks2/T2.2.11/internal/models"
	"github.com/listnt/tasks2/T2.2.11/utils"
	"github.com/sirupsen/logrus"
)

// Repository ...
type Repository interface {
	GetEvents(user_id int,date string) ([]models.Event,error)
	UpdateEvent(user_id int,date string,event string,newValue models.Event) error
	Delete(user_id int,date string,event string) error
	Store(v models.Event) error

	loadToCache() error
	Close()
}

type repository struct {
	db    *sql.DB
	cache map[int]map[string][]models.Event
	validator utils.Validator
	lg    *logrus.Logger
}

// NewRepository ...
func NewRepository(lg *logrus.Logger, cfg *Config.Config) (Repository, error) {
	rep := new(repository)
	var err error
	rep.cache = make(map[int]map[string][]models.Event, 1000)
	rep.validator=utils.ValidatorNew()
	rep.lg = lg
	rep.db, err = sql.Open("postgres", cfg.ConnectionString)
	fmt.Println(rep.db,err)
	if err != nil {
		return nil, err
	}
	if rep.loadToCache() != nil {
		rep.db.Close()
		return nil, err
	}
	return rep, nil
}
