package app

import (
	"net/http"

	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {

	v1 := e.Group("/api/v1")

	v1.GET("/task", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
