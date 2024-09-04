package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
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
		//For Web
		w.Header().Set("token", token)
		//For API
		fmt.Fprintf(w, "{token:%v}", token)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{error: Invalid Credentials}")
		return
	}
}

// Function for registering a new user (for demonstration purposes)
func Register(w http.ResponseWriter, req *http.Request) {
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{error: Invalid data}")
		return
	}

	// Remember to securely hash passwords before storing them
	user.ID = 1 // Just for demonstration purposes
	//Write user to file
	writeFile("users.csv", []string{user.Username, user.Password, strconv.Itoa(user.ID)})
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{user: %#v}", user)
}

func writeFile(filename string, user []string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	data := user
	writer.Write(data)
}
