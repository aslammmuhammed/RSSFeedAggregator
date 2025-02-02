package router

import (
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/entity"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/handler/app_health"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/handler/app_user"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/handler/feeds"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/handler/posts"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/middleware"
	"github.com/gorilla/mux"
)

func HealthRoute(router *mux.Router, apiCfg *entity.ApiCfg) {
	hh := app_health.HealthHandler{
		ApiCfg: apiCfg,
	}
	router.Handle("/healthz", hh).Methods("GET")
}

func UserRoutes(router *mux.Router, apiCfg *entity.ApiCfg) *app_user.UserHandler {
	uh := app_user.UserHandler{
		ApiCfg: apiCfg,
	}
	router.HandleFunc("/user", uh.CreateUserHandler).Methods("POST")
	router.HandleFunc("/user", middleware.UserAuthMiddleware(&uh, uh.GetUserHandler)).Methods("GET")
	return &uh
}

func FeedRoutes(router *mux.Router, apiCfg *entity.ApiCfg, uh app_user.UserHandler) *feeds.FeedHandler {
	//feeds
	fh := feeds.FeedHandler{
		ApiCfg: apiCfg,
	}
	router.HandleFunc("/feeds", middleware.UserAuthMiddleware(&uh, fh.CreateFeedHandler)).Methods("POST")
	router.HandleFunc("/feeds", fh.GetFeedsHandler).Methods("GET")
	return &fh
}

func FeedFollowRoutes(router *mux.Router, uh *app_user.UserHandler, fh *feeds.FeedHandler) {
	router.HandleFunc("/feed_follows", middleware.UserAuthMiddleware(uh, fh.CreateFeedFollowHandler)).Methods("POST")
	router.HandleFunc("/feed_follows", middleware.UserAuthMiddleware(uh, fh.GetFeedFollowsForUserHandler)).Methods("GET")
	router.HandleFunc("/feed_follows/{id}", middleware.UserAuthMiddleware(uh, fh.DeleteFeedFollowForUserHandler)).Methods("DELETE")
}

func PostRoutes(router *mux.Router, apiCfg *entity.ApiCfg, uh *app_user.UserHandler) {
	ph := posts.PostHandler{
		ApiCfg: apiCfg,
	}
	router.HandleFunc("/posts", middleware.UserAuthMiddleware(uh, ph.GetNewPostsForUser)).Methods("GET")
}
