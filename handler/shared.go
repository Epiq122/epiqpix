package handler

import (
	"log/slog"
	"net/http"

	"github.com/epiq122/epiqpixai/models"
)

func getAuthenticatedUser(r *http.Request) models.AuthenticatedUser {
	user, ok := r.Context().Value(models.UserContextKey).(models.AuthenticatedUser)
	if !ok {
		return models.AuthenticatedUser{}
	}
	return user
}

func Make(h func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("internal server error", "err", err, "path", r.URL.Path)
		}
	}
}
