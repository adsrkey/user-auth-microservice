package user

import (
	"auth-service/service/auth/internal/domain/user"
	utils "auth-service/service/auth/utils/jwt"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) Auth(ctx context.Context, user *user.User) (*utils.SignedToken, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	user.Hash = hash

	ucUser, err := uc.adapterStorage.GetUser(ctx, user)
	if err != nil {
		return nil, err
	}

	dayHours := 24
	jwtWrapper := &utils.JwtWrapper{
		SecretKey:       "jfaijfp3420",
		Issuer:          "",
		ExpirationHours: int64(dayHours),
	}

	token, err := jwtWrapper.GenerateToken(ucUser)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(ucUser.Hash, []byte(user.Password))
	if err != nil {
		return nil, err
	}

	return token, nil
}
