package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	pathToUrls := map[string]string{
		"/gophercise": "https://github.com/rohan3011/gophercises",
		"/blogs":      "http://rohankamble.me/",
	}

	mapHandler := MapHandler(pathToUrls, mux)

	http.ListenAndServe(":8080", mapHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
