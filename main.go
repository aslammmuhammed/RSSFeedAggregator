package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error reading env file: ", err)
	}

	portString, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("PORT not found in env")
	}

	host, exists := os.LookupEnv("HOST")
	if !exists {
		log.Fatal("HOST not found in env")
	}

	allowedOrigin := "http://" + host + ":" + portString
	fmt.Printf("Starting server on %v\n", allowedOrigin)

	// Main router
	muxRouter := mux.NewRouter()

	// Subrouter for v1 endpoints
	v1MuxRouter := muxRouter.PathPrefix("/v1").Subrouter()

	// Add a route to v1MuxRouter with restricted HTTP methods
	v1MuxRouter.Handle("/healthz", new(healthHandler)).Methods("GET")
	v1MuxRouter.Handle("/err", new(errorHandler)).Methods("GET")

	// CORS
	// Default CORS
	// v1MuxRouterCORS := cors.Default().Handler(v1MuxRouter)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{allowedOrigin},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		// Enable Debugging for testing, consider disabling in production
		// Debug: true,
	})

	v1MuxRouterCORS := c.Handler(v1MuxRouter)
	server := &http.Server{
		Addr:    host + ":" + portString,
		Handler: v1MuxRouterCORS, // Use the CORS-wrapped router as the handler
	}

	log.Fatal(server.ListenAndServe()) // ListenAndServe will block
}
