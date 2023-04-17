package http

import (
	"auth-service/service/auth/internal/delivery/http/response"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

func (de *Delivery) auth(echoCtx echo.Context) error {
	milliseconds := 100
	timeout := time.Millisecond * time.Duration(milliseconds)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	bearer := echoCtx.Request().Header.Get("Authorization")

	if len(bearer) == 0 {
		resp := response.ErrorResponse{
			StatusCode:       http.StatusNotAcceptable,
			DeveloperMessage: "no token",
		}
		return echoCtx.JSON(http.StatusNotAcceptable, resp)
	}

	token := strings.Split(bearer, "Bearer ")[1]

	err := de.ucUser.Auth(ctx, token)
	if err != nil {
		resp := response.ErrorResponse{
			StatusCode:       http.StatusNotAcceptable,
			DeveloperMessage: "token not valid",
		}
		return echoCtx.JSON(http.StatusNotAcceptable, resp)
	}

	resp := response.Response{Message: "token is valid"}

	return echoCtx.JSON(200, resp)
}
