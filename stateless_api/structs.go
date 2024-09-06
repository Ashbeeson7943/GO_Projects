package main

import (
	"net/http"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type KeyLimit struct {
	Last_use  time.Time
	Wait_time int
}

type Middleware func(http.Handler) http.Handler
