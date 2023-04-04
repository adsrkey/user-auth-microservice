package http

import (
	"auth-service/service/auth/internal/domain/user"
	"github.com/labstack/echo/v4"
)

func (de *Delivery) bindUser(c echo.Context) (user.User, error) {
	var u user.User
	err := c.Bind(&u)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}
