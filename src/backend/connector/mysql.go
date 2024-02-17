package connector

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func Connect() {
	// データベース接続情報を設定
	port := os.Getenv("MYSQL_PORT")
	host := os.Getenv("MYSQL_HOST")
	scheme := os.Getenv("MYSQL_SCHEME")

	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, scheme)

	// データベースに接続
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()

	// データベース接続を確認
	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")
}
