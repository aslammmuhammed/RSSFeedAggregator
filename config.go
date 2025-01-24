package main

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppPort string `yaml:"appPort" env:"APP_PORT" env-default:"3000"`
	AppHost string `yaml:"appHost" env:"APP_HOST" env-default:"localhost"`
	DBUrl   string `yaml:"dbUrl" env:"DB_URL" env-default:"postgres"`
}

// NewConfig returns app config
func NewConfig() (*Config, error) {
	cfg := &Config{}
	//loads from config file
	err := cleanenv.ReadConfig("config.yaml", cfg)
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
