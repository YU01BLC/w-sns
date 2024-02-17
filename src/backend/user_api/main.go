package main

import (
	"log"
	"net/http"

	"github.com/YU01BLC/w-sns/src/backend/connector"
	"github.com/YU01BLC/w-sns/src/backend/middleware"
	"github.com/YU01BLC/w-sns/src/backend/user_api/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	filename := "../.env"
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalln("Error loading .env")
	}
}
func main() {
	// connect databases.
	connector.Connect()
	// init echo
	e := echo.New()
	// set middleware
	logFile := middleware.Use(e)
	// Init handler
	e.GET("/", func(c echo.Context) error {
		req := c.Request()
		url := req.URL.String()
		c.Logger().Infof("Request URL: %s", url)
		return c.String(http.StatusOK, "Hello World")
	})
	router.Routing(e)

	e.Logger.Fatal(e.Start(":8080"))
	defer logFile.Close()
}
