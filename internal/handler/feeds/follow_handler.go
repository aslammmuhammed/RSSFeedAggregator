package feeds

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/entity"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/utilities"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (f *FeedHandler) CreateFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	inputparams := entity.CreateFeedFollow{}
	err := decoder.Decode(&inputparams)
	if err != nil {
		log.Printf("error decoding createfeedfollow input params: %v", err)
		utilities.RespondWithError(w, http.StatusUnprocessableEntity, "couldn't decode input params")
		return
	}
	feedFollow, err := f.ApiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    inputparams.FeedId,
	})
	if err != nil {
		log.Printf("error execuing CreateFeedFollow query: %v", err)
		utilities.RespondWithError(w, http.StatusInternalServerError, "couldn't create feed follow")
		return
	}
	log.Printf("created feed follow with id: %v ,user: %v , feed:%v", feedFollow.ID, feedFollow.UserID, feedFollow.FeedID)
	utilities.RespondWithJSON(w, http.StatusOK, utilities.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (f *FeedHandler) GetFeedFollowsForUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := f.ApiCfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		log.Printf("error executing query GetFeedFollowsForUser: %v", err)
		utilities.RespondWithError(w, http.StatusInternalServerError, "couldn't get feed follows")
		return
	}
	log.Printf("Got %d feedfollows for user %v", len(feedFollows), user.ID)
	utilities.RespondWithJSON(w, http.StatusOK, utilities.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (f *FeedHandler) DeleteFeedFollowForUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("invalid uuid in path for feed follow: %v", err)
		utilities.RespondWithError(w, http.StatusBadRequest, "invalid feed follow id")
		return
	}
	err = f.ApiCfg.DB.DeleteFeedFollowForUser(r.Context(), database.DeleteFeedFollowForUserParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		log.Printf("couldn't execute query DeleteFeedFollowForUser : %v", err)
		utilities.RespondWithError(w, http.StatusInternalServerError, "couldn't delete feed follow")
		return
	}
	log.Printf("deleting feed follow:%v for user :%v", feedFollowID, user.ID)
	utilities.RespondWithJSON(w, http.StatusOK, struct{}{})
}
