package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Pholice/blog_aggregator/internal/database"
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

	serveMux := http.NewServeMux()
	server := http.Server{
		Addr:    port,
		Handler: serveMux,
	}

	serveMux.Handle("GET /v1/healthz", http.HandlerFunc(cfg.readiness))
	serveMux.Handle("POST /v1/users", http.HandlerFunc(cfg.createUser))
	serveMux.Handle("GET /v1/err", http.HandlerFunc(cfg.err))
	server.ListenAndServe()
}
