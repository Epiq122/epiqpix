package handler

import (
	"net/http"

	"github.com/epiq122/epiqpixai/view/settings"
)

func HandleSettingsIndex(w http.ResponseWriter, r *http.Request) error {
	user := getAuthenticatedUser(r)
	return render(r, w, settings.Index(user))

}
