package storage

import (
	context "auth-service/pkg/type"
	"auth-service/service/auth/internal/domain/user"
	"auth-service/service/auth/internal/repository/storage/postgres/dao"
)

type Storage interface {
	User
}

type User interface {
	GetUser
}

type GetUser interface {
	GetUser(c context.Context, user *user.User) (*dao.User, error)
	CreateUser(c context.Context, user *user.User) error
}
