package notes

import (
	"github.com/go-martini/martini"
	"github.com/iangreenleaf/noted-go/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Notes/Server", func() {
	var server *martini.ClassicMartini
	var recorder *httptest.ResponseRecorder
	var request *http.Request

	mydb := db.NewDB(":memory:")
	dbmap := NotesMap(mydb)
	dbmap.CreateTablesIfNotExists()

	BeforeEach(func() {
		dbmap.TruncateTables()
		server = NewServer(mydb)
		recorder = httptest.NewRecorder()
	})

	Describe("#index", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/notes", nil)
			dbmap.Insert(&Note{"test note 1", "abcdefg"})
			dbmap.Insert(&Note{"test note 2", "hijk lmnop"})
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
