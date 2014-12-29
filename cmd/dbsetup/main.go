package main

import "github.com/iangreenleaf/noted-go/db"

func main() {
	mydb := db.NewDB("development")
	db.RecreateTables(mydb)
	db.Seed(mydb)
}
