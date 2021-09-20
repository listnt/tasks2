package mymodule

import "testing"

func TestCase1(t *testing.T){
	err:=WGet("https://example.com","text.txt")
	if err!=nil{
		t.Error(err)
	}
}
func TestCase2(t *testing.T){
	err:=WGet("https://example.com","")
	if err!=nil{
		t.Error(err)
	}
}
func TestCase3(t *testing.T){
	err:=WGet("wrongaddr","text.txt")
	if err==nil{
		t.Error(err)
	}
}