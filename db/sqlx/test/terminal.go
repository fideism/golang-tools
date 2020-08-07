package main

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/fideism/golang-tools/db/sqlx"
	"github.com/fideism/golang-tools/logger"
)

func main() {
	fmt.Println("main")

	logger.NewEntry(logger.Channel("sql")).Info("aaa")

	config := sqlx.Config{
		Dialect:  "postgres",
		Server:   "ip",
		Port:     5432,
		User:     "root",
		Database: "root",
		Password: "123456",
	}

	type User struct {
		ID       string `json:"id" db:"id"`
		Nickname string `json:"nickname" db:"nickname"`
	}

	var user User

	cols := []interface{}{"id", "nickname"}
	sql, args, err := sqlx.PostgresBuilder().From("user.users").Where(goqu.C("id").Eq("3d97cd16-7804-403d-9d63-d116ef17a67b")).Select(cols...).ToSQL()
	fmt.Println(sql)
	fmt.Println(args)
	fmt.Println(err)

	conn := sqlx.Build(config)

	er := conn.Get(&user, sql, args...)
	//conn, _ := db.Connect(config)
	//conn.Ping()
	//er := conn.Get(&user, sql, args...)
	fmt.Println(er)
	fmt.Println(user)
}
