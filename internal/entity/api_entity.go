package entity

import "github.com/aslammmuhammed/RSSFeedAggregator/internal/database"

type ApiCfg struct {
	DB                *database.Queries
	QueryLimit        int
	ScrapeInterval    int
	ScrapeConcurrency int
	ScrapeTimeout     int
}
