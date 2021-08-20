package mymodule

import (
	"testing"
)

var testcase1 =[]string{"Пятак", "Пятка", "Пятка", "Тяпка", "слиток", "истолк", "слиток", "столик", "слиток", "листок", "слиток", "одиночка"}
var testcase1res = map[string][]string{
	"пятак":[]string{"пятак", "пятка", "тяпка"} ,
	"слиток":[]string{"истолк", "листок", "слиток", "столик"},
}

func TestCase1(t *testing.T){
	mp:=Annagrams(&testcase1)
	for k :=range *mp{
		res:=""
		for i:=0;i<len((*mp)[k]);i++{
			res+=(*mp)[k][i]
		}
		res1:=""
		for i:=0;i<len((testcase1res)[k]);i++{
			res1+=(testcase1res)[k][i]
		}
		if res!=res1{
			t.Error("expected",res1,"got",res)
		}
	}
}