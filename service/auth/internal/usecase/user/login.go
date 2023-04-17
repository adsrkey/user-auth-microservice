package user

import (
	"auth-service/service/auth/internal/domain/user"
	utils "auth-service/service/auth/utils/jwt"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) Login(ctx context.Context, input *user.User) (*utils.SignedToken, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	input.Hash = hash

	ucUser, err := uc.adapterStorage.GetUser(ctx, input)
	if err != nil {
		return nil, err
	}

	dayHours := 1
	jwtWrapper := &utils.JwtWrapper{
		SecretKey:       uc.secretKey,
		Issuer:          ucUser.ID.String(),
		ExpirationHours: int64(dayHours),
	}

	token, err := jwtWrapper.GenerateToken(ucUser)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(ucUser.Hash, []byte(input.Password))
	if err != nil {
		return nil, err
	}

	return token, nil
}
