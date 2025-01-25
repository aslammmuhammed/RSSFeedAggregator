package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now().Local()
		log.Printf("[LoggingMiddleware] started: %s [%s] %s from %s", startTime, r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("[LoggingMiddleware] completed: %s [%s] %s from %s in %s", time.Now().Local(), r.Method, r.RequestURI, r.RemoteAddr, time.Since(startTime))
	})
}
