package mymodule

import (
	"errors"

	"github.com/beevik/ntp"
)

type Time interface {
	Time() (string, error)
	SetLayout(layout string)
	SetServer(server string)
}

// layout of time representation
const (
	defaultServer = "time.nist.gov"
	defaultLayout = "3:04:05 PM (MST) on Monday, January _2, 2006"
)

type time struct {
	timeServer string
	layout     string
}

func NewTime() Time {
	return &time{
		timeServer: defaultServer,
		layout:     defaultLayout,
	}
}

// Return
func (t *time) Time() (string, error) {
	time, err := ntp.Time(t.timeServer)
	if err != nil {
		return "", errors.New("wrong server")
	}
	return time.Local().Format(t.layout), nil
}

func (t *time) SetLayout(layout string) {
	t.layout = layout
}

func (t *time) SetServer(server string) {
	t.timeServer = server
}
