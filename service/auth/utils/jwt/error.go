package utils

import "errors"

var (
	ErrParseClaims = errors.New("couldn't parse claims")
	ErrExpired     = errors.New("jwt is expired")
)
