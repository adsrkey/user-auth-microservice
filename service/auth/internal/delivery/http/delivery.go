package http

import (
	"auth-service/service/auth/internal/usecase"
	"github.com/labstack/echo/v4"
)

type Delivery struct {
	echo   *echo.Echo
	ucUser usecase.User
}

func New(echo *echo.Echo, ucUser usecase.User) *Delivery {
	d := &Delivery{
		echo:   echo,
		ucUser: ucUser,
	}
	d.initRouter()
	return d
}
