package main

import (
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/url":
		url(w, r)
	case "/body":
		body(w, r)
	default:
		w.Write([]byte("Hello World"))
	}
}

func main() {
	http.HandleFunc("/", handleFunc)
	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}
	server.ListenAndServe()
}
