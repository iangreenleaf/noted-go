package notes

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"net/http"
)

func NewServer() *martini.ClassicMartini {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/notes", func() (int, string) {
		notes := AllNotes()
		js, err := json.Marshal(notes)
		if err != nil {
			return 500, err.Error()
		}
		return http.StatusOK, string(js)
	})

	return m
}
