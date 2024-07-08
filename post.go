package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Pholice/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) processFeed(feed database.Feed, wg *sync.WaitGroup) {
	layout := "Mon, 02 Jan 2006 15:04:05 -0700"
	defer wg.Done()
	cfg.DB.MarkFeedFetched(context.Background(), feed.ID)
	rssFeed, err := fetchXML(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:", err)
		return
	}
	for _, item := range rssFeed.Channel.Items {
		timeStr, err := time.Parse(layout, item.PubDate)
		if err != nil {
			log.Println("Time parse error")
		}
		cfg.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         sql.NullString{String: item.Link, Valid: true},
			FeedID:      feed.ID,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: timeStr},
		})
	}
}

func (cfg *apiConfig) getPost(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.DB.GetPostByUser(context.Background(),
		database.GetPostByUserParams{UserID: user.ID, Limit: 10})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not get posts")
		return
	}
	respondWithJSON(w, http.StatusOK, posts)
}
