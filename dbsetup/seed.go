package main

import "iangreenleaf/noted/db"

func main() {
	mydb := db.NewDB()
	mydb.Exec("DROP TABLE IF EXISTS notes")
	mydb.Exec("CREATE TABLE notes(title text, text text)")
	mydb.Exec("INSERT INTO notes (title, text) VALUES ('Test note', 'Just a test'), ('Buy milk', 'At the store')")
}
