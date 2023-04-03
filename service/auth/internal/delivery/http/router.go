package http

import (
	"github.com/labstack/echo/v4"
)

const (
	Api      = "/api"
	Version  = "/v1"
	AuthPath = "/auth"
	RegPath  = "/reg"
)

func (de *Delivery) initRouter() *echo.Group {
	group := de.echo.Group(Api + Version)
	group.POST(AuthPath, de.auth)
	group.POST(RegPath, de.reg)
	return nil
}
