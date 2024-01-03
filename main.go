package main

import (
	"github.com/basliq/basliq-server-mvp/config"
	"github.com/basliq/basliq-server-mvp/repository"
	"github.com/basliq/basliq-server-mvp/server/http"
	"github.com/basliq/basliq-server-mvp/service/user"
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

	repo := repository.New(repository.Config{
		Username: "db-user",
		Password: "password",
		Port:     5432,
		Host:     "localhost",
		DBName:   "basliq",
	})

	userSvc := user.New(repo)

	server := http.New(appConfig, userSvc)
	server.Serve()
}
