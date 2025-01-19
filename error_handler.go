package main

import "net/http"

type errorHandler struct {
}

func (h errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "something went wrong")
}
