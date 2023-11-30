package http

import "github.com/labstack/echo/v4"

func healthCheck(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"message": "all is ok!",
	})
}
