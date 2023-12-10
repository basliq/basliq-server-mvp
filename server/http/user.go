package http

import (
	"github.com/basliq/basliq-server-mvp/service/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) registerUser(c echo.Context) error {
	var req user.RegisterReq
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := s.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusCreated, res)
}
