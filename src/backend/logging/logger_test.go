package logging

import (
	"os"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestLoggerInit(t *testing.T) {
	type args struct {
		e *echo.Echo
	}

	e := echo.New()

	tests := []struct {
		name  string
		args  args
		level string
	}{
		{"Log test debug.", args{e}, "debug"},
		{"Log test info.", args{e}, "info"},
		{"Log test warn.", args{e}, "Warn"},
		{"Log test warning.", args{e}, "Warning"},
		{"Log test error.", args{e}, "error"},
		{"Log test off.", args{e}, "OFF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// src/backend/log/app.logを確認すること！
			os.Setenv("LOG_LEVEL", tt.level)
			LoggerInit(tt.args.e)
			logger := tt.args.e.Logger
			logger.Debug("Test debug")
			logger.Info("Test info")
			logger.Warn("test warn")
			logger.Error("test error")
		})
	}
}

func TestLoggerInitNoEnv(t *testing.T) {
	type args struct {
		e *echo.Echo
	}

	e := echo.New()

	tests := []struct {
		name string
		args args
	}{
		{"Log test debug.", args{e}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// src/backend/log/app.logを確認すること！
			// info以上が表示されるはず
			os.Unsetenv("LOG_LEVEL")
			LoggerInit(tt.args.e)
			logger := tt.args.e.Logger
			logger.Debug("Test debug")
			logger.Info("Test info")
			logger.Warn("test warn")
			logger.Error("test error")
		})
	}
}
