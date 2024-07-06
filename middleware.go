package main

import (
	"net/http"

	"github.com/Pholice/blog_aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := authToken(r)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Bad Formatted Request")
			return
		}
		user, err := cfg.DB.KeyUser(r.Context(), token)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Key Invalid")
			return
		}
		handler(w, r, user)
	}
}
