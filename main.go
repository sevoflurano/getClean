package main

import (
	"log"
	"os"
	"path/filepath"

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

func newService(envVars map[string]string) (service.Service, error) {
	// Delega a criação do Config para o arquivo do SO correspondente
	cfg := getServiceConfig(envVars)
	return service.New(&program{}, cfg)
}

func main() {
	envVars := map[string]string{}

	if len(os.Args) > 1 && os.Args[1] == "install" {
		home, err := os.UserHomeDir()
		if err == nil {
			downloadsReal := filepath.Join(home, "Downloads")
			envVars["GETCLEAN_DOWNLOADS"] = downloadsReal
		}
	}

	s, err := newService(envVars)
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
