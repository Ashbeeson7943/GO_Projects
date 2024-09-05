package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, req *http.Request) {
	var user User
	json.NewDecoder(req.Body).Decode(&user)
	// Check if credentials are valid (replace this logic with real authentication)
	if user.Username == "user" && user.Password == "password" {
		// Generate a JWT token
		token, err := GenerateJWTToken(user.ID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "{error: Error generating authentication token, msg:%v}", err)
			return
		}
		//For API
		fmt.Fprintf(w, "{token:%v}", token)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{error: Invalid Credentials}")
		return
	}
}

func Register(w http.ResponseWriter, req *http.Request) {
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{error: Invalid data}")
		return
	}

	// Todo: Hash password, generate user-Id
	// Write user to DB
	_, err = DB.Exec(InsertUser(user.Username, user.Password))
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{user: %#v}", user)
}
