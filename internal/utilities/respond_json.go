package utilities

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("responding with 5xx error : %v\n", msg)
	} else if code > 399 {
		log.Printf("responding with 4xx error : %v\n", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error encoding json:%v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
