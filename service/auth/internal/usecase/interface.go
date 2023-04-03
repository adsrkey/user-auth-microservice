package usecase

import (
	context "auth-service/pkg/type"
	"auth-service/service/auth/internal/domain/user"
	utils "auth-service/service/auth/utils/jwt"
)

type User interface {
	UserAuth
}

type UserAuth interface {
	Auth(c context.Context, user *user.User) (*utils.SignedToken, error)
	Register(c context.Context, user *user.User) error
}
