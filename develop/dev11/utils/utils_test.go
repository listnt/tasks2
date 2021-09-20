package utils

import (
	"github.com/listnt/tasks2/T2.2.11/internal/models"
	"testing"
)

func TestCase1(t *testing.T) {
	val:=ValidatorNew()
	testv :=models.Event{UserId: 2,Date:"2021-06-06",Event: "happysome",Description: "happysome"}
	if err:=val.Validate(testv);err!=nil{
		t.Error(err)
	}
}

func TestCase2(t *testing.T) {
	val:=ValidatorNew()
	testv :=models.Event{UserId: 2,Date:"2021-31-31",Event: "happysome",Description: "happysome"}
	if err:=val.Validate(testv);err==nil{
		t.Error("expected error got nothing")
	}
}

func TestCase3(t *testing.T) {
	val:=ValidatorNew()
	testv :=models.Event{UserId: -5,Date:"2021-06-06",Event: "happysome",Description: "happysome"}
	if err:=val.Validate(testv);err==nil{
		t.Error("expected error got nothing")
	}
}

func TestCase4(t *testing.T) {
	val:=ValidatorNew()
	testv :=models.Event{UserId: 2,Date:"2021-06-06",Event: "",Description: "happysome"}
	if err:=val.Validate(testv);err==nil{
		t.Error("expected error got nothing")
	}
}