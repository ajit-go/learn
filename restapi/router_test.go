package restapi

import (
	"regexp"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/json"
)

var m = minify.New()

func initMinifier() {
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
}

func handleServer(method string, path string, t *testing.T, f http.HandlerFunc, expectedStatus int) (rr *httptest.ResponseRecorder) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	//handler := http.HandlerFunc(f)

	f.ServeHTTP(rr, req)
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}

	return
}
func TestHomePage(t *testing.T) {
	rr := handleServer("GET", "/", t, HomePage, http.StatusOK)
	expected := "Hello from home"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestArticles(t *testing.T) {
	initMinifier()
	rr := handleServer("GET", "/as", t, Articles, http.StatusOK)
	expected := `{"Articles" : [{"Title":"title 1","desc":"desc 1","content":"test content"},
	{"Title":"title 2","desc":"desc 2","content":"test content"}]}`
	expected, err := m.String("text/json", expected)
	if err != nil {
		panic(err)
	}
	if strings.Trim(rr.Body.String(), " \n") != strings.Trim(expected, " \n") {
		t.Errorf("got %v want %v",
			strings.Trim(rr.Body.String(), " "), expected)
	}
}
