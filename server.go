package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"iangreenleaf/noted/models"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/notes", func() (int, string) {
		notes := models.AllNotes()
		js, err := json.Marshal(notes)
		if err != nil {
			return 500, err.Error()
		}
		return http.StatusOK, string(js)
	})
	m.Run()
}
