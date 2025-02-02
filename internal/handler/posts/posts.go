package posts

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/entity"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/utilities"
)

type PostHandler struct {
	ApiCfg *entity.ApiCfg
}

func (p *PostHandler) GetNewPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {

	limitStr := r.URL.Query().Get("limit")
	limit := p.ApiCfg.QueryLimit
	if specifedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifedLimit
	}

	posts, err := p.ApiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		log.Printf("error executing quer GetPostsForUser for user %v :%v ", user.ID, err)
		utilities.RespondWithError(w, http.StatusInternalServerError, "couldn't get posts")
		return
	}
	utilities.RespondWithJSON(w, http.StatusOK, utilities.DatabasePostsToPosts(posts))

}
