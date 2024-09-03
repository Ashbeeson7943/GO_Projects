package main

import (
	"log"
	"net/http"
)

var mux *http.ServeMux

func main() {
	mux = http.NewServeMux()
	addFunc := http.HandlerFunc(Add)
	subtractFunc := http.HandlerFunc(Subtract)
	divideFunc := http.HandlerFunc(Divide)
	multiplyFunc := http.HandlerFunc(Multiply)
	//Public Routes

	//Protected Routes
	mux.Handle("POST /add/", AuthenticationCheck(addFunc))
	mux.Handle("POST /subtract/", AuthenticationCheck(subtractFunc))
	mux.Handle("POST /multiply/", AuthenticationCheck(multiplyFunc))
	mux.Handle("POST /divide/", AuthenticationCheck(divideFunc))
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe()) // Run the http server
}
