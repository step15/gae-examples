package foo

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUrls(t *testing.T) {
	for _, ti := range []struct {
		in   string
		want string
	}{
		{"bar", "foo is 'bar'!"},
		{"hoge", "foo is 'hoge'!"},
		{"", "foo is ''!"}} {
		req, err := http.NewRequest("GET", "http://example.com/test?foo="+ti.in, nil)
		if err != nil {
			log.Fatal(err)
		}

		w := httptest.NewRecorder()
		test(w, req)
		got := w.Body.String()
		if got != ti.want {
			t.Errorf("tried /test?foo=%v got:%q want:%q", ti.in, got, ti.want)
		}
	}
}
