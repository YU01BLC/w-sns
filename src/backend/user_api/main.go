package main

import (
	"fmt"
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
	port := 8080
	host := "0.0.0.0"
	addr := fmt.Sprintf("%s:%d", host, port)

	e.Logger.Infof("Server is running at http://%s\n", addr)
	e.Logger.Fatal(e.Start(addr))
}
