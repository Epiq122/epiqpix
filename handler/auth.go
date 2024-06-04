package handler

import (
	"log/slog"
	"net/http"

	"github.com/epiq122/epiqpixai/pkg/sb"
	"github.com/epiq122/epiqpixai/pkg/util"
	"github.com/epiq122/epiqpixai/view/auth"

	"github.com/nedpals/supabase-go"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return auth.Login().Render(r.Context(), w)

}

func HandleLoginCreate(w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if !util.IsValidEmail(credentials.Email) {
		return render(r, w, auth.LoginForm(credentials, auth.LoginErrors{
			Email: "please enter a valid email address",
		}))
	}

	if reason, ok := util.ValidatePassword(credentials.Password); !ok {
		return render(r, w, auth.LoginForm(credentials, auth.LoginErrors{
			Password: reason,
		}))
	}

	resp, err := sb.Client.Auth.SignIn(r.Context(), credentials)
	if err != nil {
		slog.Error("failed to login", "err", err)
		return render(r, w, auth.LoginForm(credentials, auth.LoginErrors{
			InvalidCredentials: "please check your credentials and try again",
		}))
	}

	cookie := &http.Cookie{
		Name:     "at",
		Value:    resp.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusFound)

	return nil
}
