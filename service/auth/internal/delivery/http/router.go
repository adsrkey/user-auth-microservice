package http

import (
	"auth-service/service/auth/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

const (
	API       = "/api"
	Version   = "/v1"
	AuthPath  = "/auth"
	LoginPath = "/login"
	RegPath   = "/reg"
)

func (de *Delivery) initRouter() {
	group := de.echo.Group(API + Version)

	group.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return middleware.ServerIsUnavailableHandlerFunc(de.ctx, next)
	})

	group.POST(AuthPath, de.auth)
	group.POST(LoginPath, de.login, middleware.ContentHandlerFunc)
	group.POST(RegPath, de.register, middleware.ContentHandlerFunc)
}
