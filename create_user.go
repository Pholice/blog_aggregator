package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Pholice/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

type Request struct {
	Name string `json:"name"`
}

func (cfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	var reqBody Request
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode")
	}

	respondWithJSON(w, http.StatusOK, database.User{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      reqBody.Name,
	})
}
