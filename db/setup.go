package db

import (
	"database/sql"
)

func RecreateTables(mydb *sql.DB) {
	mydb.Exec("DROP TABLE IF EXISTS notes")
	mydb.Exec("CREATE TABLE notes(title text, text text)")
}

func Seed(mydb *sql.DB) {
	mydb.Exec("INSERT INTO notes (title, text) VALUES ('Test note', 'Just a test'), ('Buy milk', 'At the store')")
}
