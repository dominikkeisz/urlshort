package main

import (
	"fmt"
	"net/http"
	"strings"
	"urlshort/handler"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2"
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	fmt.Println("Starting the server on :8000")
	http.ListenAndServe(":8000", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}