package main

import (
	"os"

	"github.com/kardianos/service"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go Watcher()
	return nil
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "getClean",
		DisplayName: "getClean",
		Description: "Organiza arquivos da pasta Downloads",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			panic(err)
		}
		return
	}

	s.Run()
}
