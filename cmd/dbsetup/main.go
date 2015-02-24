package main

import (
	"fmt"
	"github.com/iangreenleaf/noted-go/db"
	"github.com/iangreenleaf/noted-go/notes"
)

func main() {
	mydb := db.NewDB(fmt.Sprintf("db-%s.sqlite3", "development"))
	dbmap := notes.NotesMap(mydb)
	if err := dbmap.DropTablesIfExists(); err != nil {
		panic(err)
	}
	if err := dbmap.CreateTables(); err != nil {
		panic(err)
	}
	db.Seed(mydb)
	fmt.Println("DB set up.")
}
