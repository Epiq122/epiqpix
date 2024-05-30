package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	log.Fatal(http.ListenAndServe(os.Getenv("HTTP_LISTEN_ADDR"), router))
}

func initEverything() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}
