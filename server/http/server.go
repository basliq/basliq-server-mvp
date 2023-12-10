package http

import (
	"github.com/basliq/basliq-server-mvp/config"
	"github.com/basliq/basliq-server-mvp/service/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config  config.Config
	userSvc user.Service
}

func New(config config.Config, userSvc user.Service) Server {
	return Server{config: config, userSvc: userSvc}
}

func (s Server) Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthcheck", healthCheck)

	userGroup := e.Group("/user")
	userGroup.POST("/register", s.registerUser)

	e.Logger.Fatal(e.Start(s.config.Server.Port))
}
