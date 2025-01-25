package handler

import (
	"net/http"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/utilities"
)

type HealthHandler struct {
}

func (h HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	utilities.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"status": "ok"})
}
