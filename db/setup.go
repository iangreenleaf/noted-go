package db

import (
	"database/sql"
)

func Seed(mydb *sql.DB) {
	mydb.Exec("INSERT INTO notes (title, text) VALUES ('Test note', 'Just a test'), ('Buy milk', 'At the store')")
}
