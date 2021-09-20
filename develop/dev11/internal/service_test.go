package internal

import (
	Config "github.com/listnt/tasks2/T2.2.11/config"
	"github.com/listnt/tasks2/T2.2.11/internal/models"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestCase1(t *testing.T){
	lg :=logrus.New()
	cfg := Config.ParseConfig("../config/config.json")
	s:=NewService(lg,&cfg)
	s.Close()
}
func TestCase2(t *testing.T){
	lg :=logrus.New()
	cfg := Config.ParseConfig("../config/config.json")
	s:=NewService(lg,&cfg)
	evnts,err:=s.EventForDay(3,"2021-06-06")
	if err!=nil{
		t.Error(err)
	}
	if len(evnts)==0{
		t.Error("expected something")
	}
}
func TestCase3(t *testing.T){
	lg :=logrus.New()
	cfg := Config.ParseConfig("../config/config.json")
	s:=NewService(lg,&cfg)
	evnts,err:=s.EventForDay(3,"2021-06-06")
	if err!=nil{
		t.Error(err)
	}
	if len(evnts)==0{
		t.Error("expected something")
	}
}
func TestCase4(t *testing.T){
	lg :=logrus.New()
	cfg := Config.ParseConfig("../config/config.json")
	s:=NewService(lg,&cfg)
	evnts,err:=s.EventsForWeek(3,"2021-06-06")
	if err!=nil{
		t.Error(err)
	}
	if len(evnts)==0{
		t.Error("expected something")
	}
}
func TestCase5(t *testing.T){
	lg :=logrus.New()
	cfg := Config.ParseConfig("../config/config.json")
	s:=NewService(lg,&cfg)
	evnts,err:=s.EventForMonth(3,"2021-06-06")
	if err!=nil{
		t.Error(err)
	}
	if len(evnts)==0{
		t.Error("expected something")
	}
}

func TestCase6(t *testing.T){
	lg :=logrus.New()
	cfg := Config.ParseConfig("../config/config.json")
	s:=NewService(lg,&cfg)
	err:=s.CreateEvent(models.Event{UserId: 3,Date: "2021-06-06",Event: "new happy1",Description: "test"})
	if err!=nil{
		t.Error(err)
	}
	err=s.UpdateEvent(3,"2021-06-06","new happy1",models.Event{UserId: 3,Event: "new happy1.1",Date: "2021-06-06",Description: "another b"})
	if err!=nil{
		t.Error(err)
	}
	err=s.DeleteEvent(3,"2021-06-06","new happy1.1")
	if err!=nil{
		t.Error(err)
	}
}

func TestCase7(t *testing.T){
	lg :=logrus.New()
	cfg := Config.ParseConfig("../config/config.json")
	s:=NewService(lg,&cfg)
	_,err:=s.EventForMonth(3,"2021-31-31")
	if err==nil{
		t.Error("expected error")
	}
}