package notes

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"net/http"
	"iangreenleaf/noted/db"
	"database/sql"
)

func NewServer(mydb *sql.DB) *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(db.DBHandler(mydb))
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/notes", func(db *sql.DB) (int, string) {
		notes := AllNotes(db)
		js, err := json.Marshal(notes)
		if err != nil {
			return 500, err.Error()
		}
		return http.StatusOK, string(js)
	})

	return m
}
