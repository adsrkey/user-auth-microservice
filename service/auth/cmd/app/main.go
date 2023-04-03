package main

import (
	"auth-service/pkg/store/postgres"
	"auth-service/service/auth/internal/delivery/http"
	repository "auth-service/service/auth/internal/repository/storage/postgres"
	"auth-service/service/auth/internal/usecase/user"
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

	ucUser := user.New(repo)

	e := echo.New()

	server := http.New(e, ucUser)
	server.Start(":8080")

	server.Notify()
}
