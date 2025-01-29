package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aslammmuhammed/RSSFeedAggregator/config"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/entity"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/middleware"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/router"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" //postgres driver for database/sql package
	"github.com/rs/cors"
)

func Run(appCfg *config.Config) {
	//Open a DB connection
	dbConn, err := sql.Open("postgres", appCfg.DBUrl)
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := entity.ApiCfg{
		DB: dbQueries,
	}

	allowedOrigin := "http://" + appCfg.AppHost + ":" + appCfg.AppPort
	log.Printf("Starting server on %v\n", allowedOrigin)

	// Main router
	muxRouter := mux.NewRouter()
	muxRouter.Use(middleware.LoggingMiddleware)

	// Subrouter for v1 endpoints
	v1MuxRouter := muxRouter.PathPrefix("/v1").Subrouter()

	// Add a route to v1MuxRouter with restricted HTTP methods
	router.HealthRoute(v1MuxRouter, &apiCfg)
	uh := router.UserRoutes(v1MuxRouter, &apiCfg)
	fh := router.FeedRoutes(v1MuxRouter, &apiCfg, *uh)
	router.FeedFollowRoutes(v1MuxRouter, uh, fh)

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
