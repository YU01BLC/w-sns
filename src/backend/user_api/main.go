package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// init echo
	e := echo.New()

	// Init handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
