package http

import (
	"auth-service/service/auth/internal/delivery/http/cookie"
	"auth-service/service/auth/internal/delivery/http/response"
	"auth-service/service/auth/internal/delivery/http/validator"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/labstack/echo/v4"
)

func (de *Delivery) auth(echoCtx echo.Context) error {
	milliseconds := 100
	timeout := time.Millisecond * time.Duration(milliseconds)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	user, err := de.bindUser(echoCtx)
	if err != nil {
		resp := &response.ErrorResponse{
			StatusCode:       http.StatusBadRequest,
			DeveloperMessage: "error with bind user from request",
		}

		return echoCtx.JSON(resp.StatusCode, resp)
	}

	de.echo.Validator = validator.New()

	err = validator.ValidateReqData(de.echo.Validator, user)
	if err != nil {
		resp := &response.ErrorResponse{
			StatusCode:       http.StatusBadRequest,
			DeveloperMessage: "user not valid",
		}

		return echoCtx.JSON(resp.StatusCode, resp)
	}

	token, err := de.ucUser.Auth(ctx, &user)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			resp := &response.ErrorResponse{
				StatusCode:       http.StatusBadRequest,
				DeveloperMessage: "this user is not in the database",
			}

			return echoCtx.JSON(resp.StatusCode, resp)
		}

		codesSize := 2
		codes := make([]string, 0, codesSize)

		codes = append(codes,
			ErrTerminatingConnection)

		resp, err := de.handlePgError(err, codes)
		if err != nil {
			return echoCtx.JSON(resp.StatusCode, &resp)
		}

		resp = response.ErrorResponse{
			StatusCode:       http.StatusUnauthorized,
			DeveloperMessage: "user not authorized",
		}

		return echoCtx.JSON(resp.StatusCode, &resp)
	}

	echoCtx.SetCookie(cookie.JwtTokenCookie(token))

	status := http.StatusOK
	resp := &response.Response{
		Message: "user authorized",
	}

	return echoCtx.JSON(status, resp)
}
