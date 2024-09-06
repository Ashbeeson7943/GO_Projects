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
  user_id INTEGER,
  last_use TEXT,
  wait_time INTEGER,
  FOREIGN KEY (user_id) REFERENCES users(id));`

var InsertTestKeyLimits = fmt.Sprintf(`INSERT INTO keyLimits(id, user_id, last_use, wait_time) VALUES (1, 1, '%v', 10);`, time.Now())

func InsertKeyLimit(user_id int, wait_time int) string {
	return fmt.Sprintf("INSERT INTO keyLimits(userID, last_use, wait_time) VALUES ('%v','%v','%v');", user_id, time.Now(), wait_time)
}

func GetKeyLimit(user_id int) string {
	return fmt.Sprintf("SELECT keyLimits.last_use, keyLimits.wait_time FROM keyLimits where user_id='%v';", user_id)
}

func CreateKeyForUser(username string) string {
	return fmt.Sprintf("INSERT", username)
}
