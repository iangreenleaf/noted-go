package notes

import (
	"database/sql"
	"gopkg.in/gorp.v1"
)

func NotesMap(db *sql.DB) *gorp.DbMap {
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Note{}, "notes").SetKeys(true, "Id")
	return dbmap
}

type Note struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func newNote(title string, text string) *Note {
	return &Note{0, title, text}
}

func AllNotes(db *gorp.DbMap) []Note {
	var notes []Note
	_, err := db.Select(&notes, "SELECT * FROM notes")
	if err != nil {
		panic(err)
	}

	return notes
}
