package main

import (
	"fmt"
	"iangreenleaf/noted/notes"
	"iangreenleaf/noted/db"
)

func main() {
	mydb := db.NewDB(fmt.Sprintf("db-%s.sqlite3", "development"))
  m := notes.NewServer(mydb)
	m.Run()
}
