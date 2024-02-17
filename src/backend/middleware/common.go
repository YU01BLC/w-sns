package middleware

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Use(e *echo.Echo) *os.File {
	e.Use(mid.CSRF())

	logFile := openFile()
	format := "${time_rfc3339_nano} ${host} ${uri} ${method} ${status} \n"
	e.Use(mid.LoggerWithConfig(mid.LoggerConfig{
		Format: format,
		Output: logFile,
	}))
	return logFile
}

func openFile() *os.File {
	file, ok := os.LookupEnv("LOG_FILEa")
	if !ok {
		file = "../log/access.log"
	}
	createDir(file)
	permission := 0644
	writer, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, fs.FileMode(permission))
	if err != nil {
		log.Fatalln("Do not open log file. ", file, err)
	}
	return writer
}

func createDir(filePath string) {
	dirPath := filepath.Dir(filePath)
	// ディレクトリの存在を確認
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// ディレクトリが存在しない場合は作成
		err := os.MkdirAll(dirPath, 0755) // すべての親ディレクトリも含めて作成
		if err != nil {
			// エラー処理
			log.Fatalln("Error creating directory:", err)
		}
	}
}
