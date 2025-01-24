package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

// expected API Key format is
// Authorization: APIKey xxxx-value-xxx
// space seperated value
func GetAPIKey(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		log.Println("empty auth header")
		return "", errors.New("no Authorization header provided")
	}
	authHeaderValues := strings.Split(authHeader, " ")
	if len(authHeaderValues) != 2 || authHeaderValues[0] != "APIKey" {
		return "", errors.New("malformed authorization header")
	}
	return authHeaderValues[1], nil
}
