package user

import (
	utils "auth-service/service/auth/utils/jwt"
	"context"
	"fmt"
	"time"
)

func (uc *UseCase) Auth(ctx context.Context, token string) error {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	jwtWrapper := &utils.JwtWrapper{
		SecretKey: uc.secretKey,
	}

	validateToken, err := jwtWrapper.ValidateToken(ctx, token)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", validateToken)

	return nil
}
