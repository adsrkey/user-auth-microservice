package utils

import (
	"auth-service/service/auth/internal/repository/storage/postgres/dao"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaims struct {
	jwt.StandardClaims
	SessionID uuid.UUID
	UserID    uuid.UUID
}

type SignedToken struct {
	Value     string
	ExpiresAt time.Time
}

func (w *JwtWrapper) GenerateToken(user *dao.User) (*SignedToken, error) {
	claims := &JwtClaims{
		SessionID: uuid.New(),
		UserID:    user.ID,
		StandardClaims: jwt.StandardClaims{
			Audience:  user.ID.String(),
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			//Id: ,
			IssuedAt: time.Now().Unix(),
			Issuer:   w.Issuer,
			Subject:  "logged_in",
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

func (w *JwtWrapper) ValidateToken(ctx context.Context, signedToken string) (claims *JwtClaims, err error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("context done")
	default:
	}

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
		return nil, ErrParseClaims
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, ErrExpired
	}

	return claims, nil
}
