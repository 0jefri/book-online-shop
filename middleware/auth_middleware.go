package middleware

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if token != "Bearer mysecrettoken" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		log.Println("Authorization successful")
		next.ServeHTTP(w, r)
	})
}
