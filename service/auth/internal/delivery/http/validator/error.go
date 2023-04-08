package validator

import "errors"

var (
	ErrNotValidBodyType = errors.New("request body type not valid")
)
