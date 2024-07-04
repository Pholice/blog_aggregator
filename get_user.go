package main

import (
	"net/http"
)

func (cfg *apiConfig) getUser(w http.ResponseWriter, r *http.Request) {
	token, err := getToken(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Formatted Request")
		return
	}
	row, err := cfg.DB.KeyUser(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Key Invalid")
		return
	}
	respondWithJSON(w, http.StatusOK, row)
}
