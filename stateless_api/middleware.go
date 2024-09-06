package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func AddMiddleware(next http.Handler, m ...Middleware) http.Handler {
	if len(m) < 1 {
		return next
	}
	wrapped := next
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}
	return wrapped
}

func AuthenticationCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authToken := req.Header.Get("Authorization")
		if authToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "{error: Missing authentication token}")
			return
		}
		authTokenParts := strings.Split(authToken, " ")
		if len(authTokenParts) != 2 || authTokenParts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "{error: Invalid authentication token}")
			return
		}
		authToken = authTokenParts[1]
		claims, err := VerifyJWTToken(authToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "{error: Invalid authentication token}")
			return
		}
		req.Header.Set("user_id", fmt.Sprint(claims["user_id"]))
		next.ServeHTTP(w, req)
	})
}

func RateLimitCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//Todo: Update this with the query for getting rate limits. Will need to re-write the query
		h_user_id := req.Header.Get("user_id")
		user_id, err := strconv.Atoi(h_user_id)
		if err != nil {
			log.Fatal(err)
		}
		row := DB.QueryRow(GetKeyLimit(user_id))
		var keyLimit KeyLimit //Fix: Mapping values back from query isn't mapping?
		if err = row.Scan(&keyLimit.Last_use, &keyLimit.Wait_time); err == sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{error: Error getting token limits, msg:%v}", err)
			return
		}
		log.Println(keyLimit)
		//Key rate limit logic
		log.Println("Rate Check..")
		next.ServeHTTP(w, req)
	})
}
