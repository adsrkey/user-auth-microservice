package user

import (
	"auth-service/service/auth/internal/usecase/adapters/storage"
	"os"
)

type UseCase struct {
	adapterStorage storage.User
	secretKey      string
}

func New(storage storage.User) *UseCase {

	secretKey := os.Getenv("SECRET_KEY")
	if len(secretKey) == 0 {
		err := os.Setenv("SECRET_KEY", "my_secret")
		if err != nil {
			return nil
		}
	}

	return &UseCase{
		adapterStorage: storage,
		secretKey:      secretKey,
	}
}
