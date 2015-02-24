package notes

import (
	"database/sql"
	"encoding/json"
	"github.com/go-martini/martini"
	"gopkg.in/gorp.v1"
	"net/http"
	"net/url"
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

	m.Get("/tomboy/api/1.0", func(db *gorp.DbMap, req *http.Request) (int, string) {
		makeURL := func(path string) string {
			u := url.URL{
				req.URL.Scheme,
				req.URL.Opaque,
				req.URL.User,
				req.URL.Host,
				path,
				"",
				"",
			}
			return u.String()
		}
		response := struct {
			Oauth_request_token_url string `json:"oauth_request_token_url"`
			Oauth_authorize_url     string `json:"oauth_authorize_url"`
			Oauth_access_token_url  string `json:"oauth_access_token_url"`
			ApiVersion              string `json:"api-version"`
		}{
			makeURL("/oauth/request_token"),
			makeURL("/oauth/authorize"),
			makeURL("/oauth/access_token"),
			"1.0",
		}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			return 500, err.Error()
		}
		return http.StatusOK, string(responseJSON)
	})

	return m
}
