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
	seconds := 15
	timeout := time.Millisecond * time.Duration(seconds)
	ctx, cancel := context.WithTimeout(ctx, timeout)

	defer cancel()

	log.Println("Starting shutdown auth service...")

	<-ctx.Done()

	// shutdown server
	err := de.echo.Server.Shutdown(ctx)
	if err != nil {
		return
	}

	// ...

	log.Println("Login service is shutdown...")
}
