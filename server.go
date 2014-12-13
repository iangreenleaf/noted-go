package main

import "github.com/go-martini/martini"
import "encoding/json"
import "net/http"

type Note struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/notes", func() (int, string) {
		notes := []Note{
			Note{"First note", "Get the milk"},
			Note{"Second note", "Buy bread"},
		}
		js, err := json.Marshal(notes)
		if err != nil {
			return 500, err.Error()
		}
		return http.StatusOK, string(js)
	})
	m.Run()
}
