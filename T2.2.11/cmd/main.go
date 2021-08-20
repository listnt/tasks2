package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/listnt/tasks2/T2.2.11/cmd/route"
	"github.com/sirupsen/logrus"
)

func main() {
	s := route.NewServer("./config/config.json")
	errchan := s.Launch()
	log.Println("[!!!] Starting")
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errchan:
		if err != nil {
			logrus.Error(err)
			s.Close()
			os.Exit(1)
		}
	case <-osSignals:
		log.Println("[!!!] Exiting")
		s.Close()
	}
}
