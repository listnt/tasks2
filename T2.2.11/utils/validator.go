package utils

import (
	"errors"
	"github.com/listnt/tasks2/T2.2.11/internal/models"
	"time"
)
const (
	layoutISO = "2006-01-02"
)
type Validator interface{
	Validate(event models.Event) error

}

type validator struct{}

func ValidatorNew() Validator{
	return &validator{}
}

func (v* validator) Validate(event models.Event) error{
	if event.UserId<1{
		return errors.New("validate: UserId must be more than zero")
	}
	if _,err:=time.Parse(layoutISO,event.Date);err!=nil{
		return err
	}
	if event.Event==""{
		return errors.New("validate: Event can't be empty")
	}
	return nil
}