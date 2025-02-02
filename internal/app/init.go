package app

import (
	"database/sql"
	"log"

	"github.com/aslammmuhammed/RSSFeedAggregator/config"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/entity"
)

func initAPIConfig(appCfg *config.Config) *entity.ApiCfg {
	//Open a DB connection
	dbConn, err := sql.Open("postgres", appCfg.DBUrl)
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := entity.ApiCfg{
		DB:                dbQueries,
		QueryLimit:        appCfg.DefaultQueryLimit,
		ScrapeInterval:    appCfg.ScrapeInterval,
		ScrapeConcurrency: appCfg.ScrapeConcurrency,
		ScrapeTimeout:     appCfg.ScrapeTimeout,
	}
	return &apiCfg
}
