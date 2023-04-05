package cookie

import (
	utils "auth-service/service/auth/utils/jwt"
	"net/http"
)

func JwtTokenCookie(token *utils.SignedToken) *http.Cookie {
	return &http.Cookie{
		Name:  JwtTokenName,
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
