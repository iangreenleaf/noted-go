package main

import "iangreenleaf/noted/db"

func main() {
	mydb := db.NewDB("development")
	db.RecreateTables(mydb)
	db.Seed(mydb)
}
