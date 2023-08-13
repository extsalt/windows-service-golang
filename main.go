package main

import (
	"github.com/kardianos/service"
	"log"
	"os"
	"time"
)

type birju struct{}

func (p *birju) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *birju) run() {
	for {
		log.Println("Service is running...")
		time.Sleep(5 * time.Second)
	}
}

func (p *birju) Stop(s service.Service) error {
	log.Println("Service is stopping...")
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "Birju",
		DisplayName: "Birju",
		Description: "Birju in Golang",
	}

	prg := &birju{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
