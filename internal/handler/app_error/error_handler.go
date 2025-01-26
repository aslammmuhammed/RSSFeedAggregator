package app_error

import (
	"net/http"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/utilities"
)

type ErrorHandler struct {
}

func (h ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	utilities.RespondWithError(w, http.StatusInternalServerError, "something went wrong")
}
