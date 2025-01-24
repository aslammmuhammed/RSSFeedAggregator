package main

import (
	"time"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type apiCfg struct {
	DB *database.Queries
}

type createUserParams struct {
	Name string
}
