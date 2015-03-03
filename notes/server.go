package notes

import (
	"database/sql"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
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
	m.Use(render.Renderer())
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/notes", func(db *gorp.DbMap, r render.Render) {
		notes := AllNotes(db)
		r.JSON(http.StatusOK, notes)
	})

	m.Get("/tomboy/api/1.0", func(db *gorp.DbMap, req *http.Request, r render.Render) {
		makeURL := func(path string) string {
			u := url.URL{
				req.URL.Scheme,
				"",
				req.URL.User,
				req.Host,
				path,
				"",
				"",
			}
			if u.Scheme == "" {
				u.Scheme = "http"
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
		r.JSON(http.StatusOK, response)
	})

	return m
}
