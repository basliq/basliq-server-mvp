package http

import (
	"github.com/basliq/basliq-server-mvp/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config config.Config
}

func New(config config.Config) Server {
	return Server{config: config}
}

func (s Server) Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthcheck", healthCheck)

	e.Logger.Fatal(e.Start(s.config.Server.Port))
}
