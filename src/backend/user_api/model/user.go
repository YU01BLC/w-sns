package model

import (
	"encoding/json"
	"log"

	"github.com/doug-martin/goqu/v9"
)

// 認証の時のみに使うこと
type UserPass struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// ユーザ構造体
type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(id, name, email string) *User {
	return &User{id, name, email}
}

func (u *User) Marshal() []byte {
	jsonData, err := json.Marshal(u)
	if err != nil {
		log.Println("User not Marshal. ", err)
	}
	return jsonData
}

func Unmarshal(jsonData []byte) *User {
	var u User
	err := json.Unmarshal(jsonData, &u)
	if err != nil {
		log.Println("Unmarsial. ", err)
	}
	return &u
}

func GetUsers() []User {
	sql, _, _ := goqu.From("users").ToSQL()
	log.Println(sql)
	// FIXME: DBからの取得に修正
	user1 := User{
		Id:    "12345",
		Name:  "Taro Yamada",
		Email: "taro@example.com",
	}

	user2 := User{
		Id:    "67890",
		Name:  "Hanako Sato",
		Email: "hanako@example.com",
	}
	user3 := User{
		Id:    "67891",
		Name:  "Test Sato",
		Email: "test@example.com",
	}
	users := []User{user1, user2, user3}
	return users
}
