package main

import "net/http"

func (cfg *apiConfig) readiness(w http.ResponseWriter, r *http.Request) {
	type Message struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, Message{Status: "ok"})
}

func (cfg *apiConfig) err(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
