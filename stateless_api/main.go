package main

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
	"net/http"
	"os"
)

var mux *http.ServeMux
var DB *sql.DB

func main() {
	sqllite()
	mux = http.NewServeMux()
	addFunc := http.HandlerFunc(Add)
	subtractFunc := http.HandlerFunc(Subtract)
	divideFunc := http.HandlerFunc(Divide)
	multiplyFunc := http.HandlerFunc(Multiply)
	//Public Routes
	mux.HandleFunc("POST /login", Login)
	mux.HandleFunc("POST /register", Register)
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

func sqllite() {
	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file, deletes old one
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")
	db, err := sql.Open("sqlite", "sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	DB = db
	if _, err := db.Exec(CreateUserTable); err != nil {
		log.Fatal(err.Error())
	}
}
