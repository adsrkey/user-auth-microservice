package middleware

import (
	"auth-service/service/auth/internal/delivery/http/response"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ServerIsUnavailableHandlerFunc(ctx context.Context, next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		if isUnavailable(ctx, echoCtx) {
			return nil
		}

		return next(echoCtx)
	}
}

func isUnavailable(ctx context.Context, echoCtx echo.Context) bool {
	select {
	case <-ctx.Done():
		resp := &response.ErrorResponse{
			StatusCode:       http.StatusServiceUnavailable,
			DeveloperMessage: "Service is unavailable. Server starts shutting down",
		}
		err := echoCtx.JSON(resp.StatusCode, resp)

		if err != nil {
			return true
		}

		return true
	default:
	}

	return false
}
