package main

import (
	"log"

	"github.com/aslammmuhammed/RSSFeedAggregator/config"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/app"
)

func main() {

	// Load configuration
	appCfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	// Run
	app.Run(appCfg)
}
