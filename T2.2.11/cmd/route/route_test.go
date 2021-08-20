package route

import (
	"net/http"
	"testing"
)


func TestCase1(t *testing.T){
	s:=NewServer("../../config/config.json")
	s.Launch()
	http.NewRequest("GET","/create_event/",nil)

}