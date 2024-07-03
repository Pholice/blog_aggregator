package main

import "net/http"

func (cfg *apiConfig) readiness(w http.ResponseWriter, r *http.Request) {
	type Message struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, Message{Status: "ok"})
}
