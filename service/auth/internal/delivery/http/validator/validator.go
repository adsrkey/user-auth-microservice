package validator

import (
	"auth-service/service/auth/internal/domain/user"
	"errors"
	"github.com/labstack/echo/v4"
)

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(i interface{}) error {
	switch i.(type) {
	case user.User:
		// TODO: validate email
		return nil
	default:
		return errors.New("request body type not valid")
	}
}

func ValidateReqData(validator echo.Validator, reqData user.User) error {
	if err := validator.Validate(reqData); err != nil {
		return err
	}
	return nil
}
