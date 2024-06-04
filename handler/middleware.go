package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/epiq122/epiqpixai/models"
)

const userKey = "user"

func Withuser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		user := models.AuthenticatedUser{
			Email:    "epiqpixai@gmail.com",
			LoggedIn: true,
		}
		ctx := context.WithValue(r.Context(), userKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
