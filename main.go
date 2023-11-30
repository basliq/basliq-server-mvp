package main

import (
	"github.com/basliq/basliq-server-mvp/config"
	"github.com/basliq/basliq-server-mvp/server/http"
)

// TODO - dockerize the project itself
// TODO - add postgres driver
// TODO - add jwt
// TODO - add migrator

// TODO - add validator
// TODO - add rich error

func main() {
	appConfig := config.Config{
		Server: config.ServerConfig{Port: ":8080"},
	}

	server := http.New(appConfig)
	server.Serve()
}
