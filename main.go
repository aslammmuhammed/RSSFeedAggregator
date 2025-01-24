package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //
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

	dbUrl, exists := os.LookupEnv("DB_URL")
	if !exists {
		log.Fatal("HOST not found in env")
	}

	//Open a DB connection
	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := apiCfg{
		DB: dbQueries,
	}

	allowedOrigin := "http://" + host + ":" + portString
	log.Printf("Starting server on %v\n", allowedOrigin)

	// Main router
	muxRouter := mux.NewRouter()
	muxRouter.Use(LoggingMiddleware)

	// Subrouter for v1 endpoints
	v1MuxRouter := muxRouter.PathPrefix("/v1").Subrouter()

	// Add a route to v1MuxRouter with restricted HTTP methods
	v1MuxRouter.Handle("/healthz", new(healthHandler)).Methods("GET")
	v1MuxRouter.Handle("/err", new(errorHandler)).Methods("GET")
	v1MuxRouter.HandleFunc("/user", apiCfg.createUserHandler).Methods("POST")
	v1MuxRouter.HandleFunc("/user", apiCfg.getUserHandler).Methods("GET")

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{allowedOrigin},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	// CORS Wrapped on muxRouter
	muxRoutersCORSWrappedHandler := c.Handler(muxRouter)
	server := &http.Server{
		Addr:    host + ":" + portString,
		Handler: muxRoutersCORSWrappedHandler, // Use the CORS-wrapped router as the handler
	}

	log.Fatal(server.ListenAndServe()) // ListenAndServe will block
}
