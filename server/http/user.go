package http

import (
	"github.com/basliq/basliq-server-mvp/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) signupUser(c echo.Context) error {
	var req dto.SignupReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	res, err := s.userSvc.Signup(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusCreated, res)
}

func (s Server) loginUser(c echo.Context) error {
	var req dto.LoginReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	res, err := s.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, res)
}
