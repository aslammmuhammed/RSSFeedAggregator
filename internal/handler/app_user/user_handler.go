package app_user

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/entity"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/utilities"
	"github.com/google/uuid"
)

type UserHandler struct {
	ApiCfg *entity.ApiCfg
}

func (u *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	inputParams := entity.CreateUserParams{}
	err := decoder.Decode(&inputParams)
	if err != nil {
		log.Printf("Error decoding create user inputParams: %v", err)
		utilities.RespondWithError(w, http.StatusUnprocessableEntity, "couldn't decode input parameters")
		return
	}

	user, err := u.ApiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:      inputParams.Name,
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		log.Printf("Error executing create user query: %v", err)
		utilities.RespondWithError(w, http.StatusInternalServerError, "couldnot create user")
		return
	}

	log.Printf("user %v created\n", user.Name)
	utilities.RespondWithJSON(w, http.StatusOK, utilities.DatabaseUserToUser(user))
}

func (u *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	utilities.RespondWithJSON(w, http.StatusOK, utilities.DatabaseUserToUser(user))
}
