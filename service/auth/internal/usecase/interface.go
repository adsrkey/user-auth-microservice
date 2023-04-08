package usecase

import (
	"auth-service/service/auth/internal/domain/user"
	utils "auth-service/service/auth/utils/jwt"
	"context"
)

type User interface {
	UserAuth
}

type UserAuth interface {
	Auth(ctx context.Context, user *user.User) (*utils.SignedToken, error)
	Register(ctx context.Context, user *user.User) error
}
