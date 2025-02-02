package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type ScrapeConfig struct {
	Concurrency int
}
type Config struct {
	AppPort           string `env-required:"true" yaml:"appPort" env:"APP_PORT" env-default:"3000"`
	AppHost           string `env-required:"true" yaml:"appHost" env:"APP_HOST" env-default:"localhost"`
	DBUrl             string `env-required:"true" yaml:"dbUrl" env:"DB_URL" env-default:"postgres"`
	DefaultQueryLimit int    `env-required:"true" yaml:"defaultQueryLimit" env:"DEFAULT_QUERY_LIMIT" env-default:"10"`
	ScrapeInterval    int    `env-required:"true" yaml:"scrapeInterval" env:"SCRAPE_INTERVAL" env-default:"10"`
	ScrapeConcurrency int    `env-required:"true" yaml:"scrapeConcurrency" env:"SCRAPE_CONCURRENCY" env-default:"10"`
	ScrapeTimeout     int    `env-required:"true" yaml:"scrapeTimeout" env:"SCRAPE_TIMEOUT" env-default:"10"`
}

// NewConfig returns app config
func NewConfig() (*Config, error) {
	cfg := &Config{}
	//loads from config file
	err := cleanenv.ReadConfig("config/config.yaml", cfg)
	if err != nil {
		log.Printf("error Reading config file: %v", err)
		return nil, fmt.Errorf("config error: %w", err)
	}

	//Environment variables
	//Overrides values from file
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
