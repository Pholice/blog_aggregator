package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Pholice/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type Request struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	var reqBody Request
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode JSON request")
		return
	}
	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      reqBody.Name,
		Url:       reqBody.URL,
		UserID:    user.ID,
	}
	feed, err := cfg.DB.CreateFeed(r.Context(), newFeed)
	if err != nil {
		log.Printf("DB error: %v", err)
		respondWithError(w, http.StatusBadRequest, "Could not save to DB")
		return
	}
	respondWithJSON(w, http.StatusOK, feed)
}
