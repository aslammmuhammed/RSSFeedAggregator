package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/auth"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/google/uuid"
)

func (a *apiCfg) createUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	inputParams := createUserParams{}
	err := decoder.Decode(&inputParams)
	if err != nil {
		log.Printf("Error decoding create user inputParams: %v", err)
		respondWithError(w, http.StatusUnprocessableEntity, "couldn't decode input parameters")
		return
	}

	user, err := a.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:      inputParams.Name,
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		log.Printf("Error executing create user query: %v", err)
		respondWithError(w, http.StatusInternalServerError, "couldnot create user")
		return
	}

	log.Printf("user %v created\n", user.Name)
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (a *apiCfg) getUserHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		log.Printf("couldn't find api key: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
		return
	}
	user, err := a.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		log.Printf("no user found with api key: %v\n", err)
		respondWithError(w, http.StatusNotFound, "Couldn't get user")
		return
	}
	log.Printf("user with API key found :%v", user)
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
