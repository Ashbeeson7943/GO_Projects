package main

import (
	"log"
	"net/http"
)

var mux *http.ServeMux

func main() {
	mux = http.NewServeMux()
	mux.HandleFunc("POST /add/", Add)
	mux.HandleFunc("POST /subtract/", Subtract)
	mux.HandleFunc("POST /multiply/", Multiply)
	mux.HandleFunc("POST /divide/", Divide)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe()) // Run the http server
}
