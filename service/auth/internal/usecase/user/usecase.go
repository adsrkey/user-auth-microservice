package user

import (
	"auth-service/service/auth/internal/usecase/adapters/storage"
)

type UseCase struct {
	adapterStorage storage.User
}

func New(storage storage.User) *UseCase {
	return &UseCase{
		adapterStorage: storage,
	}
}
