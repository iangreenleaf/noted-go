package notes

import (
	"github.com/go-martini/martini"
	"github.com/iangreenleaf/noted-go/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
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
		var note1 *Note
		var note2 *Note

		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/notes", nil)
			note1 = &Note{0, "test note 1", "abcdefg"}
			dbmap.Insert(note1)
			note2 = &Note{0, "test note 2", "hijk lmnop"}
			dbmap.Insert(note2)
		})

		It("is successful", func() {
			server.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(200))
		})

		It("returns JSON", func() {
			server.ServeHTTP(recorder, request)
			expected := fmt.Sprintf(`[
				{ "id": %d, "title": "test note 1", "text": "abcdefg" },
				{ "id": %d, "title": "test note 2", "text": "hijk lmnop" }
			]`, note1.Id, note2.Id)
			Expect(recorder.Body.Bytes()).To(MatchJSON(expected))
		})
	})
})
