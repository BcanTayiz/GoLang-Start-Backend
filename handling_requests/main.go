package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/ninja":
		fmt.Fprint(w, "Wallace")
	default:
		fmt.Fprint(w, "Big Fat Error")
	}
}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlVsPlain")
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Attempt")
	w.Header().Set("Content-Type", "text/html")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func main() {
	http.HandleFunc("/", htmlVsPlain)
	http.HandleFunc("/timeout", timeout)
	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}
	server.ListenAndServe()
}
