package notes

import "database/sql"

type Note struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func AllNotes(db *sql.DB) []Note {
	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		panic(err)
	}

	var title, text string
	result := make([]Note, 0, 100)
	for rows.Next() {
		rows.Scan(&title, &text)
		result = append(result, Note{title, text})
	}
	return result
}