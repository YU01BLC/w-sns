package controller

import (
	"net/http"

	"github.com/YU01BLC/w-sns/src/backend/user_api/model"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	log := c.Logger()
	users := model.GetUsers()
	log.Debugf("Get users: %v", users)
	return c.JSON(http.StatusOK, users)
}
