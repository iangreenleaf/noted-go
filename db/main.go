package db

import (
	"database/sql"
	"github.com/go-martini/martini"
	_ "github.com/mattn/go-sqlite3"
)

func NewDB(name string) *sql.DB {
	db, err := sql.Open("sqlite3", name)
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
