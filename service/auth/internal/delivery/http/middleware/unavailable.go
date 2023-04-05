package middleware

import (
	"auth-service/service/auth/internal/delivery/http/response"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ServerIsUnavailableHandlerFunc(ctx context.Context, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if isUnavailable(ctx, c) {
			return nil
		}
		return next(c)
	}
}

func isUnavailable(ctx context.Context, c echo.Context) bool {
	select {
	case <-ctx.Done():
		resp := &response.ErrorResponse{
			StatusCode:       http.StatusServiceUnavailable,
			DeveloperMessage: "Service is unavailable. Server starts shutting down",
		}
		err := c.JSON(resp.StatusCode, resp)
		if err != nil {
			return true
		}
		return true
	default:
	}
	return false
}
