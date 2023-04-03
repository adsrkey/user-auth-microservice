package user

import (
	"auth-service/service/auth/internal/usecase/adapters/storage"
	utils "auth-service/service/auth/utils/jwt"
)

type UseCase struct {
	adapterStorage storage.User
	jwt            *utils.JwtWrapper
}

func New(storage storage.User, jwt *utils.JwtWrapper) *UseCase {
	return &UseCase{
		adapterStorage: storage,
		jwt:            jwt,
	}
}
