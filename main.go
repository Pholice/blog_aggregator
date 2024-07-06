package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Pholice/blog_aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("dbURL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Could not connect to DB")
	}
	dbQueries := database.New(db)
	cfg := apiConfig{DB: dbQueries}

	v1Router := chi.NewRouter()
	server := http.Server{
		Addr:    port,
		Handler: v1Router,
	}

	v1Router.Get("/v1/healthz", http.HandlerFunc(cfg.readiness))
	v1Router.Get("/v1/err", http.HandlerFunc(cfg.err))
	v1Router.Get("/v1/users", cfg.middlewareAuth(cfg.getUser))
	v1Router.Get("/v1/feeds", http.HandlerFunc(cfg.getFeed))
	v1Router.Get("/v1/feed_follows", cfg.middlewareAuth(cfg.getFeedFollow))
	v1Router.Post("/v1/users", http.HandlerFunc(cfg.createUser))
	v1Router.Post("/v1/feeds", cfg.middlewareAuth(cfg.createFeed))
	v1Router.Post("/v1/feed_follows", cfg.middlewareAuth(cfg.createFeedFollow))
	v1Router.Delete("/v1/feed_follows/{feedFollowID}", http.HandlerFunc(cfg.deleteFeedFollow))
	server.ListenAndServe()
}
