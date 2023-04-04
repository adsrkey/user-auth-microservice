package http

import (
	context "auth-service/pkg/type"
	"auth-service/service/auth/internal/delivery/http/validator"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (de *Delivery) reg(c echo.Context) error {
	ctx := context.New(c)
	ctx.WithTimeout(100 * time.Millisecond)

	user, err := de.bindUser(c)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err = validator.ValidateReqData(de.echo, user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err = de.ucUser.Register(ctx, &user)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			const ErrCodeDuplicateUniqueConstraint = "23505"
			const ErrTerminatingConnection = "57P01"

			if pgErr.Code == ErrCodeDuplicateUniqueConstraint {
				return c.String(http.StatusConflict, "user with such data is already registered")
			}

			if pgErr.Code == ErrTerminatingConnection {
				return c.String(http.StatusServiceUnavailable, "user not authorized")
			}

			return c.String(http.StatusInternalServerError, "create user error")
		}

		return c.String(http.StatusInternalServerError, "user not registered")
	}
	return c.String(http.StatusCreated, "user registered")
}
