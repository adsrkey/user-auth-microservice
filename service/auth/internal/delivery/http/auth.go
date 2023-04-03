package http

import (
	context "auth-service/pkg/type"
	validator "auth-service/service/auth/internal/delivery/http/validator/auth"
	"auth-service/service/auth/internal/domain/user"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (de *Delivery) auth(c echo.Context) error {
	var timeout = 300 * time.Millisecond
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

	token, err := de.ucUser.Auth(ctx, &reqData)
	if err != nil {
		return c.String(http.StatusUnauthorized, "user not authorized")
	}

	c.SetCookie(jwtTokenCookie(token))

	return c.String(http.StatusAccepted, "user authorized")
}
