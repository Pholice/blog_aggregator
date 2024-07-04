package main

import (
	"encoding/json"
	"log"
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
		respondWithError(w, http.StatusBadRequest, "Could not decode JSON request")
		return
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      reqBody.Name,
	}
	user, err := cfg.DB.CreateUser(r.Context(), newUser)
	if err != nil {
		log.Printf("DB error: %v", err)
		respondWithError(w, http.StatusBadRequest, "Could not save to DB")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}
