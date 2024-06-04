package view

import (
	"context"

	"github.com/epiq122/epiqpixai/models"
)

func AuthenticatedUser(ctx context.Context) models.AuthenticatedUser {
	user, ok := ctx.Value(models.UserContextKey).(models.AuthenticatedUser)
	if !ok {
		return models.AuthenticatedUser{}
	}
	return user

}
