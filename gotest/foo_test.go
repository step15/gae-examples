package foo

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testUrls(t *testing.T) {
	for _, test := range []struct {
		in   string
		want string
	}{
		{"bar", "foo is 'bar'"},
		{"hoge", "foo is 'hoge'"},
		{"", "foo is ''"}} {
		req, err := http.NewRequest("GET", "http://example.com/test", nil)
		if err != nil {
			log.Fatal(err)
		}

		w := httptest.NewRecorder()
		root(w, req)
		got := w.Body.String()
		if got != test.want {
			t.Errorf("tried /test?foo=%v. got: %v want %v", test.in, got, test.want)
		}
	}
}
