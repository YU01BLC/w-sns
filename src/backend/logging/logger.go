package logging

import (
	"io"
	"os"
	"strings"

	"github.com/YU01BLC/w-sns/src/backend/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func LoggerInit(e *echo.Echo) {
	logLevelStr, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevelStr = "INFO"
	}

	logLevelStr = strings.ToUpper(logLevelStr)

	if l, ok := e.Logger.(*log.Logger); ok {
		setLevel(logLevelStr, l)
		l.SetOutput(openFile())
		l.SetPrefix("${time_rfc3339_nano} ${file} ${line} [${level}]")
	}
}

func setLevel(level string, l *log.Logger) {
	switch level {
	case "DEBUG":
		l.SetLevel(log.DEBUG)
	case "INFO":
		l.SetLevel(log.INFO)
	case "WARNING":
		l.SetLevel(log.WARN)
	case "WARN":
		l.SetLevel(log.WARN)
	case "ERROR":
		l.SetLevel(log.ERROR)
	case "OFF":
		l.SetLevel(log.OFF)
	}
}

func openFile() io.Writer {
	filePath, ok := os.LookupEnv("APP_LOG_FILE")
	if !ok {
		filePath = "../log/app.log"
	}
	util.CreateDir(filePath)
	permission := 0644
	writer := util.Openfile(filePath, permission)
	mw := io.MultiWriter(os.Stdout, writer)
	return mw
}
