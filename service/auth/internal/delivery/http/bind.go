package http

import (
	"auth-service/service/auth/internal/domain/user"

	"github.com/labstack/echo/v4"
)

func (de *Delivery) bindUser(c echo.Context) (user.User, error) {
	var userBind user.User

	err := c.Bind(&userBind)
	if err != nil {
		return user.User{}, err
	}

	return userBind, nil
}
