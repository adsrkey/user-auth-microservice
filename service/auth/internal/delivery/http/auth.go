package http

import (
	context "auth-service/pkg/type"
	"auth-service/service/auth/internal/delivery/http/validator"
	"auth-service/service/auth/internal/usecase"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func (de *Delivery) auth(c echo.Context) error {
	ctx := context.New(c)
	ctx.WithTimeout(300 * time.Millisecond)

	user, err := de.bindUser(c)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err = validator.ValidateReqData(de.echo, user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	token, err := de.ucUser.Auth(ctx, &user)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, usecase.ErrorSidNotFound) {
			return c.String(http.StatusBadRequest, "bad request")
		}
		if errors.As(err, &pgErr) {
			const ErrTerminatingConnection = "57P01"

			if pgErr.Code == ErrTerminatingConnection {
				return c.String(http.StatusServiceUnavailable, "user not authorized")
			}

			log.Println(err)
			return c.String(http.StatusInternalServerError, "user not authorized")
		}
		return c.String(http.StatusUnauthorized, "user not authorized")
	}

	c.SetCookie(jwtTokenCookie(token))

	return c.String(http.StatusAccepted, "user authorized")
}
