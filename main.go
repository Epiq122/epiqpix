package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/epiq122/epiqpixai/handler"
	"github.com/epiq122/epiqpixai/pkg/sb"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	router.Use(handler.WithUser)

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.Make(handler.HandleHomeIndex))
	router.Get("/login", handler.Make(handler.HandleLoginIndex))
	router.Get("/signup", handler.Make(handler.HandleSignupIndex))
	router.Post("/signup", handler.Make(handler.HandleSignupCreate))
	router.Post("/login", handler.Make(handler.HandleLoginCreate))
	router.Post("/logout", handler.Make(handler.HandleLogoutCreate))
	router.Get("/auth/callback", handler.Make(handler.HandleAuthCallBack))

	// protected routes
	router.Group(func(auth chi.Router) {
		auth.Use(handler.WithAuth)

		auth.Get("/settings", handler.Make(handler.HandleSettingsIndex))
	})

	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("application running", "port", port)
	log.Fatal(http.ListenAndServe(os.Getenv("HTTP_LISTEN_ADDR"), router))
}

func initEverything() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return sb.Init()
}
