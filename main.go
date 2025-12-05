package main

import (
	"net/http"
)

// http.Handle("/", http.FileServer(http.Dir("/tmp")))
func main() {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))

	srv.ListenAndServe()
}
