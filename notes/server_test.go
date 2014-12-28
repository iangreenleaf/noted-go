package notes

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/go-martini/martini"
	"iangreenleaf/noted/db"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Notes/Server", func() {
	var server *martini.ClassicMartini
	var recorder *httptest.ResponseRecorder
	var request *http.Request

	mydb := db.NewDB(":memory:")

	BeforeEach(func() {
		db.RecreateTables(mydb)
		server = NewServer(mydb)
		recorder = httptest.NewRecorder()
	})

	Describe("#index", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/notes", nil)
			mydb.Exec("INSERT INTO notes (title, text) VALUES ('test note 1', 'abcdefg'), ('test note 2', 'hijk lmnop')")
		})

		It("is successful", func() {
			server.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(200))
		})

		It("returns JSON", func() {
			server.ServeHTTP(recorder, request)
			expected := `[
				{ "title": "test note 1", "text": "abcdefg" },
				{ "title": "test note 2", "text": "hijk lmnop" }
			]`
			Expect(recorder.Body.Bytes()).To(MatchJSON(expected))
		})
	})
})
