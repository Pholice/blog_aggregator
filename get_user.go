package main

import (
	"net/http"

	"github.com/Pholice/blog_aggregator/internal/database"
)

func (cfg *apiConfig) getUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}
