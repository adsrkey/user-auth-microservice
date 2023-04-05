package http

import (
	"auth-service/service/auth/internal/usecase"
	"context"
	"github.com/labstack/echo/v4"
)

type Delivery struct {
	ctx context.Context

	echo   *echo.Echo
	ucUser usecase.User
}

func New(ctx context.Context, echo *echo.Echo, ucUser usecase.User) *Delivery {
	d := &Delivery{
		ctx:    ctx,
		echo:   echo,
		ucUser: ucUser,
	}
	d.initRouter()
	return d
}
