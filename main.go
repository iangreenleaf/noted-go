package main

import (
	"fmt"
	"github.com/iangreenleaf/noted-go/notes"
	"github.com/iangreenleaf/noted-go/db"
)

func main() {
	mydb := db.NewDB(fmt.Sprintf("db-%s.sqlite3", "development"))
  m := notes.NewServer(mydb)
	m.Run()
}
