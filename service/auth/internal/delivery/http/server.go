package http

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (de *Delivery) Start(address string) {
	go func() {
		err := de.echo.Start(address)
		if err != nil {
			log.Println(err)
		}
	}()
}

func (de *Delivery) Notify(sigint chan os.Signal) {
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
	<-sigint
}

func (de *Delivery) Shutdown(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	log.Println("Starting shutdown auth service...")

	<-ctx.Done()

	// shutdown server
	err := de.echo.Server.Shutdown(ctx)
	if err != nil {
		return
	}

	// ...

	log.Println("Auth service is shutdown...")
}
