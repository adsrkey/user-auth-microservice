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
	Login(ctx context.Context, user *user.User) (*utils.SignedToken, error)
	Auth(ctx context.Context, token string) error
	Register(ctx context.Context, user *user.User) error
}
