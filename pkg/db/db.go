package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDb() {
	var err error
	connStr := "user=postgres password=postgres dbname=prueba port=5433 sslmode=disable"

	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("failed to connect to db", err)
	}

	println("db connected")
}
