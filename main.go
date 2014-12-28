package main

import (
	"iangreenleaf/noted/notes"
	"iangreenleaf/noted/db"
)

func main() {
	mydb := db.NewDB("development")
  m := notes.NewServer(mydb)
	m.Run()
}
