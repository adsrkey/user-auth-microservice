package http

import (
	"auth-service/service/auth/internal/delivery/http/response"
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgconn"
)

const (
	ErrTerminatingConnection         = "57P01"
	ErrCodeDuplicateUniqueConstraint = "23505"
)

func (de *Delivery) initErrorResponses() {
	de.errorResponses = make(map[string]response.ErrorResponse)

	de.errorResponses[ErrCodeDuplicateUniqueConstraint] = response.ErrorResponse{
		StatusCode:       http.StatusConflict,
		DeveloperMessage: "user with such data is already registered",
	}

	de.errorResponses[ErrTerminatingConnection] = response.ErrorResponse{
		StatusCode:       http.StatusServiceUnavailable,
		DeveloperMessage: "database is unavailable",
	}
}

func (de *Delivery) handlePgError(err error, codes []string) (response.ErrorResponse, error) {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		for _, code := range codes {
			if pgErr.Code == code {
				return de.errorResponses[code], err
			}
		}

		resp := response.ErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			DeveloperMessage: "error with database",
		}

		return resp, err
	}

	if errors.Is(err, context.DeadlineExceeded) {
		resp := response.ErrorResponse{
			StatusCode:       http.StatusServiceUnavailable,
			DeveloperMessage: "context deadline exceeded",
		}

		return resp, err
	}

	return response.ErrorResponse{}, nil
}
