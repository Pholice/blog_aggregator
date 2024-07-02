package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	serveMux := http.NewServeMux()
	server := http.Server{
		Addr:    port,
		Handler: serveMux,
	}

	serveMux.Handle("GET /v1/healthz", http.HandlerFunc(readiness))
	serveMux.Handle("GET /v1/err", http.HandlerFunc(err))
	server.ListenAndServe()
}
