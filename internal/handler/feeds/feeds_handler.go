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
)

type FeedHandler struct {
	ApiCfg *entity.ApiCfg
}

func (f *FeedHandler) CreateFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	inputParams := entity.CreateFeedParams{}
	err := decoder.Decode(&inputParams)
	if err != nil {
		log.Printf("couldn't decode create feed inputparams: %v", err)
		utilities.RespondWithError(w, http.StatusUnprocessableEntity, "couldn't decode inputparams")
		return
	}
	feed, err := f.ApiCfg.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      inputParams.Name,
		Url:       inputParams.Url,
		UserID:    user.ID,
	})
	if err != nil {
		log.Printf("Error executing create user query: %v", err)
		utilities.RespondWithError(w, http.StatusInternalServerError, "couldn't create feed")
		return
	}
	log.Printf("created feed with Name: %v,Url:%v", feed.Name, feed.Url)
	utilities.RespondWithJSON(w, http.StatusOK, utilities.DatabaseFeedToFeed(feed))
}
