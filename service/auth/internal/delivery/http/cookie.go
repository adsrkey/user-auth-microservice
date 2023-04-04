package http

import (
	"auth-service/service/auth/internal/delivery/http/cookie"
	utils "auth-service/service/auth/utils/jwt"
	"net/http"
)

func jwtTokenCookie(token *utils.SignedToken) *http.Cookie {
	return &http.Cookie{
		Name:  cookie.JwtTokenName,
		Value: token.Value,
		Path:  "/",
		//Domain:     "",
		Expires: token.ExpiresAt,
		//RawExpires: "",
		//MaxAge:     0,
		Secure:   false,
		HttpOnly: true,
		//SameSite:   0,
		//Raw:        "",
		//Unparsed:   nil,
	}
}
