package http

import (
	context "auth-service/pkg/type"
	validator "auth-service/service/auth/internal/delivery/http/validator/auth"
	"auth-service/service/auth/internal/domain/user"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (de *Delivery) reg(c echo.Context) error {
	var timeout = 1 * time.Second

	ctx := context.New(c)
	ctx.WithTimeout(timeout)

	var reqData user.User
	err := c.Bind(&reqData)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	de.echo.Validator = validator.New()
	if err = de.echo.Validator.Validate(reqData); err != nil {
		return err
	}

	err = de.ucUser.Register(ctx, &reqData)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			const ErrCodeDuplicateUniqueConstraint = "23505"

			if pgErr.Code == ErrCodeDuplicateUniqueConstraint {
				return c.String(http.StatusConflict, "user with such data is already registered")
			}

			return c.String(http.StatusInternalServerError, "create user error")
		}

		return c.String(http.StatusConflict, "user not registered")
	}

	return c.String(http.StatusCreated, "user registered")
}
