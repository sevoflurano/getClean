package main

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go func() {
		if err := WatchDownloads(); err != nil {
			log.Println(err)
		}
	}()
	return nil
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func newService() (service.Service, error) {
	cfg := &service.Config{
		Name:        "getClean",
		DisplayName: "getClean",
		Description: "Organiza arquivos da pasta Downloads",
	}

	return service.New(&program{}, cfg)
}

func main() {
	s, err := newService()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		if err := service.Control(s, os.Args[1]); err != nil {
			log.Fatal(err)
		}
		return
	}

	if service.Interactive() {
		if err := WatchDownloads(); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
