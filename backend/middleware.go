package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func LoggerMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request details
		log.Printf("Request: %s %s %s\n", r.Method, r.URL, r.Proto)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return handlers.CORS(
		// handlers.AllowedOrigins([]string{"*"}), // You can specify allowed origins here
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowCredentials(),
	)(next)
}
