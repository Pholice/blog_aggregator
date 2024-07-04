package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not marshal data")
	}
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type payload struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, payload{Error: msg})
}

func getToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	tokenString := ""
	if strings.HasPrefix(header, "ApiKey ") {
		tokenString = header[7:]
	}
	return tokenString, nil
}
