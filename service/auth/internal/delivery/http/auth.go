package http

import (
	context "auth-service/pkg/type"
	"auth-service/service/auth/internal/delivery/http/cookie"
	"auth-service/service/auth/internal/delivery/http/response"
	"auth-service/service/auth/internal/delivery/http/validator"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (de *Delivery) auth(c echo.Context) error {
	ctx := context.New(c)
	ctx.WithTimeout(100 * time.Millisecond)
	defer ctx.Cancel()

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

	token, err := de.ucUser.Auth(ctx, &user)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			resp := &response.ErrorResponse{
				StatusCode:       http.StatusBadRequest,
				DeveloperMessage: "this user is not in the database",
			}
			return c.JSON(resp.StatusCode, resp)
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			if pgErr.Code == ErrTerminatingConnection {
				resp := &response.ErrorResponse{
					StatusCode:       http.StatusServiceUnavailable,
					DeveloperMessage: "database is unavailable",
				}
				return c.JSON(resp.StatusCode, resp)
			}

			resp := &response.ErrorResponse{
				StatusCode:       http.StatusInternalServerError,
				DeveloperMessage: "error with database",
			}
			return c.JSON(resp.StatusCode, resp)
		}

		if errors.Is(err, context.DeadlineExceeded) {
			resp := &response.ErrorResponse{
				StatusCode:       http.StatusServiceUnavailable,
				DeveloperMessage: "context deadline exceeded",
			}
			return c.JSON(resp.StatusCode, resp)
		}

		resp := &response.ErrorResponse{
			StatusCode:       http.StatusUnauthorized,
			DeveloperMessage: "user not authorized",
		}
		return c.JSON(resp.StatusCode, resp)
	}

	c.SetCookie(cookie.JwtTokenCookie(token))

	status := http.StatusAccepted
	resp := &response.Response{
		Message: "user authorized",
	}
	return c.JSON(status, resp)
}
