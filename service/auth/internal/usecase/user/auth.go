package user

import (
	context "auth-service/pkg/type"
	"auth-service/service/auth/internal/domain/user"
	utils "auth-service/service/auth/utils/jwt"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (uc *UseCase) Auth(c context.Context, user *user.User) (*utils.SignedToken, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	user.Hash = hash

	ucUser, err := uc.adapterStorage.GetUser(c, user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	jwtWrapper := &utils.JwtWrapper{
		SecretKey:       "key",
		Issuer:          "",
		ExpirationHours: 24,
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
