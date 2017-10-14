package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type Controller struct {
}

func (ctl Controller) rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func test1() {
	http.HandleFunc("/", Controller.rootHandler)
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)

	log.Fatal(http.ListenAndServe(":18080", nil))
}

func main() {
	test1()
}
