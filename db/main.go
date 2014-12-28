package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-martini/martini"
)

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	return db
}

func DBHandler(db *sql.DB) martini.Handler {
	return func(context martini.Context) {
		context.Map(db)
		context.Next()
	}
}
