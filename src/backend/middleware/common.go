package middleware

import (
	"io/fs"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Use(e *echo.Echo) *os.File {
	e.Use(mid.CSRF())

	logFile := openFile()
	e.Use(mid.LoggerWithConfig(mid.LoggerConfig{
		Output: logFile,
	}))
	return logFile
}

func openFile() *os.File {
	file, ok := os.LookupEnv("LOG_FILE")
	if !ok {
		file = "../log/access.log"
	}
	permission := 0644
	writer, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, fs.FileMode(permission))
	if err != nil {
		log.Fatalln("Do not open log file. ", file, err)
	}
	return writer
}
