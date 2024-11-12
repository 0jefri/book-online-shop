package middleware

import (
	"log"
	"net/http"
	"runtime"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stack := make([]byte, 1024)
				runtime.Stack(stack, false)
				log.Printf("Recovered from panic: %v\n%s", err, stack)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
