package worker

import (
	"auth-service/pkg/store/postgres"
	"context"
	"errors"
	"log"
	"os"
	"syscall"
	"time"
)

type Worker struct {
	ctx    context.Context
	conn   *postgres.Store
	sigint chan os.Signal
}

func New(ctx context.Context, conn *postgres.Store, sigint chan os.Signal) *Worker {
	return &Worker{
		ctx:    ctx,
		conn:   conn,
		sigint: sigint,
	}
}

func (w *Worker) Run() {
	go func() {
		for {
			select {
			case <-w.ctx.Done():
				return
			case <-time.After(5 * time.Second):
				log.Println("Ping postgres connection")
				w.ping()
			}
		}
	}()
}

func (w *Worker) ping() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if err := w.conn.Pool.Ping(ctx); err != nil {
		log.Println(err)
		// TODO: reconnect to postgres or read new postgres envs from config
		err = w.reconnect()
		if err != nil {
			log.Println(err)
			log.Println("context tries connect timeout, terminate auth service immediately with graceful shutdown")
			w.sigint <- syscall.SIGTERM
		}
	}
}

func (w *Worker) reconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return errors.New("no connection")
		case <-time.After(100 * time.Millisecond):
			var err error
			log.Println("Try to reconnect to postgres")
			w.conn, err = postgres.New(postgres.Settings{})
			if err != nil {
				continue
			}
			log.Printf("Connection with postgres established, max connection size: %v \n", w.conn.Pool.Stat().MaxConns())
			return nil
		}
	}
}
