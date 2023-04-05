package http

import (
	context "auth-service/pkg/type"
	"auth-service/service/auth/internal/delivery/http/response"
	"auth-service/service/auth/internal/delivery/http/validator"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (de *Delivery) register(c echo.Context) error {
	ctx := context.New(c)
	ctx.WithTimeout(100 * time.Millisecond)

	user, err := de.bindUser(c)
	if err != nil {
		resp := &response.ErrorResponse{
			StatusCode:       http.StatusBadRequest,
			DeveloperMessage: "error with bind user from request",
		}
		return c.JSON(resp.StatusCode, resp)
	}

	de.echo.Validator = validator.New()
	err = validator.ValidateReqData(de.echo.Validator, user)
	if err != nil {
		resp := &response.ErrorResponse{
			StatusCode:       http.StatusBadRequest,
			DeveloperMessage: "user not valid",
		}
		return c.JSON(resp.StatusCode, resp)
	}

	err = de.ucUser.Register(ctx, &user)
	if err != nil {
		resp := handlePgError(err)
		return c.JSON(resp.StatusCode, &resp)
	}

	status := http.StatusCreated
	resp := &response.Response{
		Message: "user registered",
	}

	return c.JSON(status, resp)
}
