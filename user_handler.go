package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/google/uuid"
)

func (a *apiCfg) userHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	inputParams := createUserParams{}
	err := decoder.Decode(&inputParams)
	if err != nil {
		fmt.Printf("Error decoding create user inputParams: %v", err)
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
		fmt.Printf("Error executing create user query: %v", err)
		respondWithError(w, http.StatusInternalServerError, "couldnot create user")
		return
	}

	fmt.Printf("user %v created\n", user.Name)
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
