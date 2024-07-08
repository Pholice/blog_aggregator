package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
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

func authToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	tokenString := ""
	if strings.HasPrefix(header, "ApiKey ") {
		tokenString = header[7:]
	}
	return tokenString, nil
}

func (cfg *apiConfig) scraper() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		feeds, err := cfg.DB.GetNextFeedToFetch(context.Background(), 10)
		if err != nil {
			log.Println("Error fetching feeds:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go cfg.processFeed(feed, wg)
		}

		wg.Wait()
	}
}
