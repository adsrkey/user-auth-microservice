package auth

import (
	"auth-service/service/auth/internal/domain/user"
	"errors"
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
