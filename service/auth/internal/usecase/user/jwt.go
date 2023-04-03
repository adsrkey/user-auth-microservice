package user

import utils "auth-service/service/auth/utils/jwt"

func (uc *UseCase) JwtWrapper() *utils.JwtWrapper {
	return uc.jwt
}
