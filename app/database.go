package app

import (
	"database/sql"
	"go-clean/helper"
	"log"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	// * Connect to Postgres
	db, err := sql.Open("postgres", "postgres://postgres:218799@localhost:5432/go-restful-api?sslmode=disable")
	helper.NewPanicError(err)
	log.Println("Connected into database")
	return db
}
