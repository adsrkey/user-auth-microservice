package middleware

import (
	"auth-service/service/auth/internal/delivery/http/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ContentHandlerFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("Content-Type") != "application/json" {
			c.Response().Header().Set(echo.HeaderAccept, "application/json")
			resp := &response.ErrorResponse{
				StatusCode:       http.StatusUnsupportedMediaType,
				DeveloperMessage: "Unsupported Content-Type. Please set Content-Type: application/json",
			}
			return c.JSON(resp.StatusCode, resp)
		}

		if c.Request().Header.Get("Content-Length") == "0" {
			resp := &response.ErrorResponse{
				StatusCode:       http.StatusLengthRequired,
				DeveloperMessage: "Content-Length required",
			}
			return c.JSON(resp.StatusCode, resp)
		}
		return next(c)
	}
}
