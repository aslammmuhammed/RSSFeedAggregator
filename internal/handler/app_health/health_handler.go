package app_health

import (
	"log"
	"net/http"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/entity"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/utilities"
)

type HealthHandler struct {
	ApiCfg *entity.ApiCfg
}

func (h HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := h.ApiCfg.DB.CheckDatabaseHealth(r.Context())
	if err != nil {
		//no need to panic exit, since the orchestrator should kill the app on non OK response
		log.Printf("error executing health check  query: %v", err)
		utilities.RespondWithError(w, http.StatusInternalServerError, "couldn't connect to database")
		return
	}
	utilities.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"status": "ok"})
}
