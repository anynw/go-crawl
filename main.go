package main

import "net/http"

type helloHandler struct {
}

func (hh *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

type aboutHandler struct {
}

func (hh *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World"))
	// })

	// The handler is typically nil, in which case the DefaultServeMux is used.
	// http.ListenAndServe("localhost:8080", nil) // DefaultSeverMux

	hh := helloHandler{}
	ab := aboutHandler{}
	server := http.Server{
		Addr: "localhost:8080",
		// Handler: &myHandler{},
		Handler: nil,
	}

	http.Handle("/hello", &hh)
	http.Handle("/about", &ab)

	server.ListenAndServe()
}
