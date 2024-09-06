package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "modernc.org/sqlite"
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
	mux.Handle("POST /add/", AddMiddleware(addFunc, AuthenticationCheck, RateLimitCheck))
	mux.Handle("POST /subtract/", AddMiddleware(subtractFunc, AuthenticationCheck, RateLimitCheck))
	mux.Handle("POST /multiply/", AddMiddleware(multiplyFunc, AuthenticationCheck, RateLimitCheck))
	mux.Handle("POST /divide/", AddMiddleware(divideFunc, AuthenticationCheck, RateLimitCheck))
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
	log.Println("Create User Table")
	if _, err := db.Exec(CreateUserTable); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Create key limit Table")
	if _, err := db.Exec(CreateKeyLimitTable); err != nil {
		log.Fatal(err.Error())
	}
	seedDB()
}

func seedDB() {
	log.Println("Adding test users...")
	_, err := DB.Exec(InsertTestUsers)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Test users added")
	log.Println("Adding test keylimits...")
	_, err = DB.Exec(InsertTestKeyLimits)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Test keyLimits added")
}
