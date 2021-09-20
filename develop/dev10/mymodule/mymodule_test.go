package mymodule

import (
	"testing"
)

func TestCase1(t *testing.T){
	Telnet(Myflags{Timeout: 3,Host: "localhost",Port:"22"})
}