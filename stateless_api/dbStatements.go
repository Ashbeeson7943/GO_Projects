package main

import (
	"fmt"
	"time"
)

const CreateUserTable = `CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY,
  username TEXT NOT NULL,
  password TEXT NOT NULL);`

const InsertTestUsers = `INSERT INTO users(username, password) VALUES ('test', 'test');`

func InsertUser(username string, password string) string {
	return fmt.Sprintf("INSERT INTO users(username, password) VALUES ('%v', '%v');", username, password)
}

func GetUser(username string) string {
	return fmt.Sprintf("SELECT * FROM users WHERE username='%v';", username)
}

const CreateKeyLimitTable = `CREATE TABLE IF NOT EXISTS keyLimits (
  id INTEGER NOT NULL PRIMARY KEY,
  FOREIGN KEY(user_id) REFERENCES users(id),
  last_use TEXT,
  wait_time INTEGER);`

func InsertKeyLimit(user_id int, wait_time int) string {
	return fmt.Sprintf("INSERT INTO keyLimits(userID, last_use, wait_time) VALUES ('%v','%v','%v');", user_id, time.Now(), wait_time)
}

func GetKeyLimit(username string) string {
	return fmt.Sprintf("SELECT keyLimits.last_use, keyLimits.wait_time FROM keyLimits INNER JOIN users ON keyLimit.user_id=users.id where users.username='%v';", username)
}
