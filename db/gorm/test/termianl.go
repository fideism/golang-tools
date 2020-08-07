package main

import (
	"fmt"
	"github.com/fideism/golang-tools/db/gorm"
)

// User 。。。
type User struct {
	ID       string `json:"id" db:"id"`
	Nickname string `json:"nickname" db:"nickname"`
}

// TableName 返回结构体对应数据库表名
func (u *User) TableName() string {
	return "user.users"
}

func main() {
	config := gorm.Config{
		Dialect:  "postgres",
		Server:   "ip",
		Port:     5432,
		User:     "root",
		Database: "root",
		Password: "root",
	}

	conn := gorm.Build(config)

	var user User

	if err := conn.First(&user, "id = ?", "3d97cd16-7804-403d-9d63-d116ef17a67b").Error; nil != err {
		fmt.Println(err)
	}

	fmt.Println(user)

	sqlite()
}

func sqlite() {
	config := gorm.Config{
		Dialect: gorm.ConfigDialectSqliteV3,
		Server:  "path/sqlite.db",
		Database: "sqlite",
	}

	conn := gorm.Build(config)

	fmt.Println(conn)
}
