package middleware

import (
	"auth-service/service/auth/internal/delivery/http/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ContentHandlerFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Request().Header.Get("Content-Type") != "application/json" {
			ctx.Response().Header().Set(echo.HeaderAccept, "application/json")

			resp := &response.ErrorResponse{
				StatusCode:       http.StatusUnsupportedMediaType,
				DeveloperMessage: "Unsupported Content-Type. Please set Content-Type: application/json",
			}

			return ctx.JSON(resp.StatusCode, resp)
		}

		if ctx.Request().Header.Get("Content-Length") == "0" {
			resp := &response.ErrorResponse{
				StatusCode:       http.StatusLengthRequired,
				DeveloperMessage: "Content-Length required",
			}

			return ctx.JSON(resp.StatusCode, resp)
		}

		return next(ctx)
	}
}
