package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db := dbOpen()
	SelectAll(db)
	Select(db, "0001")

	UpdateUser(db, "taro", "0002")
}

func dbOpen() *sql.DB {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=user dbname=default password=secret sslmode=disable")
	if nil != err {
		log.Fatal(err)
	}

	return db
}
