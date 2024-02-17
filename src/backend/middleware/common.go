package middleware

import (
	"io"
	"os"

	"github.com/YU01BLC/w-sns/src/backend/util"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Use(e *echo.Echo) {
	e.Use(mid.CSRF())

	logFile := openFile()
	format := "${time_rfc3339_nano} ${host} ${uri} ${method} ${status} \n"
	e.Use(mid.LoggerWithConfig(mid.LoggerConfig{
		Format: format,
		Output: logFile,
	}))
}

func openFile() io.Writer {
	file, ok := os.LookupEnv("LOG_FILEa")
	if !ok {
		file = "../log/access.log"
	}
	util.CreateDir(file)
	permission := 0644
	writer := util.Openfile(file, permission)
	mw := io.MultiWriter(os.Stdout, writer)
	return mw
}
