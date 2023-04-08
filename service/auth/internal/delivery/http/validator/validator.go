package validator

import (
	"auth-service/service/auth/internal/domain/user"
	"github.com/labstack/echo/v4"
)

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(i interface{}) error {
	switch i.(type) {
	case user.User:
		return nil
	default:
		return ErrNotValidBodyType
	}
}

func ValidateReqData(validator echo.Validator, reqData user.User) error {
	err := validator.Validate(reqData)
	if err != nil {
		return err
	}

	return nil
}
