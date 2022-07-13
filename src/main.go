package main

import (
	"net/http"

	"github.com/YoshitakaSS/golang-temp/common/app"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	app.Route(e)

	e.Start(":8080")
}
