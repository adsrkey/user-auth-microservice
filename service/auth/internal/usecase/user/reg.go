package user

import (
	"auth-service/service/auth/internal/domain/user"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) Register(ctx context.Context, user *user.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)

		return err
	}

	user.Hash = hash

	err = uc.adapterStorage.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
