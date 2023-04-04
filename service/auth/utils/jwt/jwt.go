package utils

import (
	"auth-service/service/auth/internal/repository/storage/postgres/dao"
	"crypto/rand"
	"errors"
	"github.com/google/uuid"
	"io"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaims struct {
	jwt.StandardClaims
	ID    uuid.UUID
	Email string
}

type SignedToken struct {
	Value     string
	ExpiresAt time.Time
}

func (w *JwtWrapper) GenerateToken(user *dao.User) (*SignedToken, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return nil, err
	}
	claims := &JwtClaims{
		ID:    uuid.NewSHA1(user.ID, b),
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return nil, err
	}

	return &SignedToken{
		Value:     signedToken,
		ExpiresAt: time.Unix(claims.ExpiresAt, 0),
	}, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("jwt is expired")
	}

	return claims, nil
}
