package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Pholice/blog_aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (cfg *apiConfig) createFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type Request struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	var reqBody Request
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode JSON request")
		return
	}
	feed, err := cfg.DB.GetFeedID(r.Context(), reqBody.FeedID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not get feed")
		return
	}
	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), newFeedFollow)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create feed follow")
		return
	}
	respondWithJSON(w, http.StatusOK, feedFollow)
}

func (cfg *apiConfig) deleteFeedFollow(w http.ResponseWriter, r *http.Request) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not parse str to uuid")
		return
	}
	_, err = cfg.DB.DeleteFeedFollow(r.Context(), feedFollowID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could delete feed follow")
		return
	}
	respondWithJSON(w, http.StatusOK, nil)
}

func (cfg *apiConfig) getFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := cfg.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not get feed follow")
		return
	}
	respondWithJSON(w, http.StatusOK, feeds)
}
