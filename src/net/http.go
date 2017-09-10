package main

import (
	"fmt"
	"net/http"
	"html"
	"log"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func test1() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)
	
	log.Fatal(http.ListenAndServe(":18080", nil))
}

func main() {
	test1()
}
