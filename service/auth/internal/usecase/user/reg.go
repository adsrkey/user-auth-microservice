package user

import (
	context "auth-service/pkg/type"
	"auth-service/service/auth/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (uc *UseCase) Register(c context.Context, user *user.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return err
	}
	user.Hash = hash

	err = uc.adapterStorage.CreateUser(c, user)
	if err != nil {
		return err
	}
	return nil
}
