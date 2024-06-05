package handler

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
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

func hxRedirect(w http.ResponseWriter, r *http.Request, to string) error {
	if len(r.Header.Get("HX-Request")) > 0 {
		w.Header().Set("HX-Redirect", to)
		w.WriteHeader(http.StatusFound)
		return nil
	}
	http.Redirect(w, r, to, http.StatusFound)
	return nil
}

func render(r *http.Request, w http.ResponseWriter, component templ.Component) error {

	return component.Render(r.Context(), w)
}
