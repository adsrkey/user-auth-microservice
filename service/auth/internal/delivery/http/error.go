package http

import (
	context "auth-service/pkg/type"
	"auth-service/service/auth/internal/delivery/http/response"
	"errors"
	"github.com/jackc/pgconn"
	"net/http"
)

// postgres code
const (
	ErrTerminatingConnection         = "57P01"
	ErrCodeDuplicateUniqueConstraint = "23505"
)

func handlePgError(err error) response.ErrorResponse {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {

		if pgErr.Code == ErrCodeDuplicateUniqueConstraint {
			resp := response.ErrorResponse{
				StatusCode:       http.StatusConflict,
				DeveloperMessage: "user with such data is already registered",
			}
			return resp
		}

		if pgErr.Code == ErrTerminatingConnection {
			resp := response.ErrorResponse{
				StatusCode:       http.StatusServiceUnavailable,
				DeveloperMessage: "database is unavailable",
			}
			return resp
		}

		resp := response.ErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			DeveloperMessage: "error with database",
		}
		return resp
	}

	if errors.Is(err, context.DeadlineExceeded) {
		resp := response.ErrorResponse{
			StatusCode:       http.StatusServiceUnavailable,
			DeveloperMessage: "context deadline exceeded",
		}
		return resp
	}

	resp := response.ErrorResponse{
		StatusCode:       http.StatusUnauthorized,
		DeveloperMessage: "user not registered",
	}
	return resp
}
