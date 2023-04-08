package storage

import (
	"auth-service/service/auth/internal/domain/user"
	"auth-service/service/auth/internal/repository/storage/postgres/dao"
	"context"
)

type Storage interface {
	User
}

type User interface {
	GetUser
}

type GetUser interface {
	GetUser(ctx context.Context, user *user.User) (*dao.User, error)
	CreateUser(ctx context.Context, user *user.User) error
}
