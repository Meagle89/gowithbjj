package main

import (
	"log"
	"net/http"
)

func LoggerMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request details
		log.Printf("Request: %s %s %s\n", r.Method, r.URL, r.Proto)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
