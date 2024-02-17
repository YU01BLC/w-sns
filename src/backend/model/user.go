package model

import (
	"log"

	"github.com/doug-martin/goqu/v9"
)

func GetUsers() {
	sql, _, _ := goqu.From("users").ToSQL()
	log.Println(sql)
}
