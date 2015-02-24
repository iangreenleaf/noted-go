package notes

import (
	"database/sql"
	"encoding/json"
	"github.com/go-martini/martini"
	"gopkg.in/gorp.v1"
	"net/http"
)

func NotesMapHandler(db *sql.DB) martini.Handler {
	return func(context martini.Context) {
		context.Map(NotesMap(db))
		context.Next()
	}
}

func NewServer(mydb *sql.DB) *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(NotesMapHandler(mydb))
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/notes", func(db *gorp.DbMap) (int, string) {
		notes := AllNotes(db)
		js, err := json.Marshal(notes)
		if err != nil {
			return 500, err.Error()
		}
		return http.StatusOK, string(js)
	})

	return m
}
