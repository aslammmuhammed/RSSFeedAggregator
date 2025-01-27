package middleware

import (
	"log"
	"net/http"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/auth"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/handler/app_user"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/utilities"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func UserAuthMiddleware(uh *app_user.UserHandler, handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			log.Printf("[UserAuthMiddleware] couldn't find api key: %v", err)
			utilities.RespondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}
		user, err := uh.ApiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			log.Printf("[UserAuthMiddleware] no user found with api key: %v\n", err)
			utilities.RespondWithError(w, http.StatusNotFound, "Couldn't get user")
			return
		}
		log.Printf("[UserAuthMiddleware] user with API key found id: %v name: %v", user.ID, user.Name)
		handler(w, r, user)
	}
}
