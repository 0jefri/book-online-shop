package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware mencatat log dengan level INFO
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mulai waktu
		start := time.Now()

		// Proses request
		next.ServeHTTP(w, r)

		// Catat durasi, URL, dan method setelah request selesai
		duration := time.Since(start)
		log.Printf("[INFO] %s %s took %s", r.Method, r.URL.Path, duration)
	})
}
