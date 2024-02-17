package util

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func CreateDir(filePath string) {
	dirPath := filepath.Dir(filePath)
	// ディレクトリの存在を確認
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// ディレクトリが存在しない場合は作成
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			// エラー処理
			log.Fatalln("Error creating directory:", err)
		}
	}
}

func Openfile(file string, permission int) *os.File {
	writer, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, fs.FileMode(permission))
	if err != nil {
		log.Fatalln("Do not open log file. ", file, err)
	}
	return writer
}
