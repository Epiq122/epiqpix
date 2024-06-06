package db

import (
	"context"

	"github.com/epiq122/epiqpixai/models"
	"github.com/google/uuid"
)

func GetAccountbyUserID(userID uuid.UUID) (models.Account, error) {
	var account models.Account
	err := Bun.NewSelect().Model(&account).Where("user_id = ?", userID).Scan(context.Background())
	return account, err
}

func CreateAccount(account *models.Account) error {
	_, err := Bun.NewInsert().Model(account).Exec(context.Background())
	return err
}
