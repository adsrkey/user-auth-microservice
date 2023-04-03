package http

import (
	"os"
	"os/signal"
	"syscall"
)

func (de *Delivery) Start(address string) {
	go func() {
		err := de.echo.Start(address)
		if err != nil {
			panic(err)
		}
	}()
}

func (de *Delivery) Notify() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
	<-sigint
}
