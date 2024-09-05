package main

import "fmt"

const CreateUserTable = `CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY,
  username TEXT NOT NULL,
  password TEXT NOT NULL);`

func InsertUser(username string, password string) string {
	return fmt.Sprintf("INSERT INTO users(username, password) VALUES ('%v', '%v');", username, password)
}
