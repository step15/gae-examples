package foo

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/test", test)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Hello step class!")
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "foo is '%s'!", r.FormValue("foo"))
}
