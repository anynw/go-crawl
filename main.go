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

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
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

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	})

	// http.HandleFunc("/welcome", welcome)
	http.Handle("/welcome", http.HandlerFunc(welcome))

	server.ListenAndServe()
}
