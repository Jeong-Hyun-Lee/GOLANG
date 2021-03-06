package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello Foo!!</h1>")
}
func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "localhost/bar?name=\"message!\""
	}
	fmt.Fprintf(w, "%s", name)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello world!!</h1>")
}

func main() {
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/", homeHandler)
	http.Handle("/foo", &fooHandler{})
	http.ListenAndServe(":8080", nil)
}
