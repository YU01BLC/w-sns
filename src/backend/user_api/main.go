package main

import (
	"net/http"

	"github.com/YU01BLC/w-sns/src/backend/connector"
	"github.com/YU01BLC/w-sns/src/backend/logging"
	"github.com/YU01BLC/w-sns/src/backend/middleware"
	"github.com/YU01BLC/w-sns/src/backend/user_api/router"
	"github.com/YU01BLC/w-sns/src/backend/util"
	"github.com/labstack/echo/v4"
)

func init() {
	filename := "../.env"
	util.OpenEnv(filename)
}
func main() {
	// connect databases.
	connector.Connect()
	// init echo
	e := echo.New()
	// set middleware
	middleware.Use(e)
	// init app logger
	logging.LoggerInit(e)
	// Init handler
	e.GET("/", func(c echo.Context) error {
		req := c.Request()
		url := req.URL.String()
		c.Logger().Infof("Request URL: %s", url)
		return c.String(http.StatusOK, "Hello World")
	})
	router.Routing(e)

	e.Logger.Fatal(e.Start(":8080"))
	defer connector.Db.Close()
}
