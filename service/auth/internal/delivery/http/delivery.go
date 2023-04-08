package http

import (
	"auth-service/service/auth/internal/delivery/http/response"
	"auth-service/service/auth/internal/usecase"
	"context"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	ctx context.Context

	echo   *echo.Echo
	ucUser usecase.User

	errorResponses map[string]response.ErrorResponse
}

func New(ctx context.Context, echo *echo.Echo, ucUser usecase.User) *Delivery {
	delivery := &Delivery{
		ctx:    ctx,
		echo:   echo,
		ucUser: ucUser,
	}
	delivery.initRouter()
	delivery.initErrorResponses()

	return delivery
}
