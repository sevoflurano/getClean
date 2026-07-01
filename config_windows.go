//go:build windows

package main

import "github.com/kardianos/service"

func getServiceConfig(envVars map[string]string) *service.Config {
	return &service.Config{
		Name:        "getClean",
		DisplayName: "getClean",
		Description: "Organiza arquivos da pasta Downloads",
		EnvVars:     envVars,
		Option: service.KeyValue{
			"StartType": "automatic",
		},
	}
}
