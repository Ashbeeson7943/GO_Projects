package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, req *http.Request) {
	var requested_user User
	json.NewDecoder(req.Body).Decode(&requested_user)
	db_user, err := GetUserFromDB(requested_user.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{error: username not found, msg:%v}", err)
		return
	}

	if requested_user.Username == db_user.Username && requested_user.Password == db_user.Password {
		// Generate a JWT token
		token, err := GenerateJWTToken(db_user.ID)
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
	var requested_user User
	err := json.NewDecoder(req.Body).Decode(&requested_user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{error: Invalid data}")
		return
	}

	// Todo: Hash password, generate user-Id
	// Write user to DB
	_, err = DB.Exec(InsertUser(requested_user.Username, requested_user.Password))
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{user: %#v}", requested_user)
}

func GetUserFromDB(username string) (User, error) {
	row := DB.QueryRow(GetUser(username))
	var db_user User
	var err error
	if err = row.Scan(&db_user.ID, &db_user.Username, &db_user.Password); err == sql.ErrNoRows {
		return User{}, err
	}
	return db_user, nil
}
