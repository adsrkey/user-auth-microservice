package main

import (
	"auth-service/pkg/store/postgres"
	"auth-service/service/auth/internal/delivery/http"
	repository "auth-service/service/auth/internal/repository/storage/postgres"
	"auth-service/service/auth/internal/repository/storage/postgres/worker"
	"auth-service/service/auth/internal/usecase/user"
	"context"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
	_ "github.com/spf13/viper/remote"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigint := make(chan os.Signal, 1)

	// connection to db
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	defer conn.Pool.Close()

	// start reconnection worker
	pause := 5
	workerPause := time.Second * time.Duration(pause)

	timeout := 15
	reconnectTimeout := time.Second * time.Duration(timeout)

	reconnectionWorker := worker.New(
		ctx,
		conn,
		sigint,
		workerPause,
		reconnectTimeout)

	reconnectionWorker.Run()

	// repository
	repo, err := repository.New(conn.Pool, repository.Options{})
	if err != nil {
		log.Println(err)

		return
	}

	// use case
	ucUser := user.New(repo)

	// echo framework
	e := echo.New()
	e.Use(middleware.RequestID())

	// start server
	server := http.New(ctx, e, ucUser)
	server.Start(":8080")

	// wait for SIGINT/SIGTERM syscall
	server.Notify(sigint)

	// cancel reconnection workers context
	cancel()

	// graceful shutdown service
	server.Shutdown(context.Background())
}
