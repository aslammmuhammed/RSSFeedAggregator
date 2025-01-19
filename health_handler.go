package main

import "net/http"

type healthHandler struct {
}

func (h healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"status": "ok"})
}
