package cookie

import (
	utils "auth-service/service/auth/utils/jwt"
	"net/http"
)

func JwtTokenCookie(token *utils.SignedToken) *http.Cookie {
	return &http.Cookie{
		Name:     JwtTokenName,
		Value:    token.Value,
		Path:     "/",
		Expires:  token.ExpiresAt,
		Secure:   true,
		HttpOnly: true,
	}
}
