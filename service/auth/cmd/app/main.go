package main

import (
	"auth-service/pkg/store/postgres"
	"auth-service/service/auth/internal/delivery/http"
	repository "auth-service/service/auth/internal/repository/storage/postgres"
	"auth-service/service/auth/internal/usecase/user"
	utils "auth-service/service/auth/utils/jwt"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}

	repo, err := repository.New(conn.Pool, repository.Options{})
	if err != nil {
		log.Println(err)
		return
	}

	jwtWrapper := &utils.JwtWrapper{
		SecretKey:       "key",
		Issuer:          "",
		ExpirationHours: 24,
	}
	ucUser := user.New(repo, jwtWrapper)

	e := echo.New()

	server := http.New(e, ucUser)
	server.Start(":8080")

	server.Notify()
}
