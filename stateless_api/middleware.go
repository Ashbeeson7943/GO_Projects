package main

import (
	"fmt"
	"net/http"
	"strings"
)

func AuthenticationCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
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
		_, err := VerifyJWTToken(authToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "{error: Invalid authentication token}")
			return
		}
		//Not needed atm - errors when used
		//r.Header.Set("user_id", claims["user_id"].(string))
		next.ServeHTTP(w, r)
	})
}
