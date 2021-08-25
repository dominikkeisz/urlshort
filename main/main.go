package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"urlshort/handler"
)

func main() {
	mux := defaultMux()
	fptr := flag.String("fpath", "example.yaml", "file path to read yaml from")
	flag.Parse()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yaml, err := ioutil.ReadFile(*fptr)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8000")
	http.ListenAndServe(":8000", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
