package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" //
	"github.com/rs/cors"
)

func main() {
	// Load environment variables
	appCfg, err := NewConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}
	//Open a DB connection
	dbConn, err := sql.Open("postgres", appCfg.DBUrl)
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := apiCfg{
		DB: dbQueries,
	}

	allowedOrigin := "http://" + appCfg.AppHost + ":" + appCfg.AppPort
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
		Addr:    appCfg.AppHost + ":" + appCfg.AppPort,
		Handler: muxRoutersCORSWrappedHandler, // Use the CORS-wrapped router as the handler
	}

	log.Fatal(server.ListenAndServe()) // ListenAndServe will block
}
