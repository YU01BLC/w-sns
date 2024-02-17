package router

import (
	"github.com/YU01BLC/w-sns/src/backend/user_api/controller"
	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) {
	e.GET("/user", controller.GetUser)
}
