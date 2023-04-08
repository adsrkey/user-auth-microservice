package http

import (
	"auth-service/service/auth/internal/delivery/http/response"
	"auth-service/service/auth/internal/delivery/http/validator"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (de *Delivery) register(echoCtx echo.Context) error {
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

	err = de.ucUser.Register(ctx, &user)
	if err != nil {
		codesSize := 2
		codes := make([]string, 0, codesSize)

		codes = append(codes,
			ErrTerminatingConnection,
			ErrCodeDuplicateUniqueConstraint)

		resp, err := de.handlePgError(err, codes)
		if err != nil {
			return echoCtx.JSON(resp.StatusCode, &resp)
		}

		resp = response.ErrorResponse{
			StatusCode:       http.StatusUnauthorized,
			DeveloperMessage: "user not registered",
		}

		return echoCtx.JSON(resp.StatusCode, &resp)
	}

	status := http.StatusCreated
	resp := &response.Response{
		Message: "user registered",
	}

	return echoCtx.JSON(status, resp)
}
