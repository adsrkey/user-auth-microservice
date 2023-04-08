package worker

import (
	"auth-service/pkg/store/postgres"
	"context"
	"log"
	"os"
	"syscall"
	"time"
)

type Worker struct {
	ctx              context.Context
	conn             *postgres.Store
	sigint           chan os.Signal
	pause            time.Duration
	reconnectTimeout time.Duration
}

func New(ctx context.Context,
	conn *postgres.Store,
	sigint chan os.Signal,
	pause time.Duration,
	reconnectTimeout time.Duration) *Worker {
	return &Worker{
		ctx:              ctx,
		conn:             conn,
		sigint:           sigint,
		pause:            pause,
		reconnectTimeout: reconnectTimeout,
	}
}

func (w *Worker) Run() {
	go func() {
		for {
			select {
			case <-w.ctx.Done():
				return
			case <-time.After(w.pause):
				log.Println("Ping postgres connection")
				w.ping()
			}
		}
	}()
}

func (w *Worker) ping() {
	ctx, cancel := context.WithTimeout(context.Background(), w.reconnectTimeout)
	defer cancel()

	if err := w.conn.Pool.Ping(ctx); err != nil {
		log.Println(err)
		err = w.reconnect()

		if err != nil {
			log.Println(err)
			log.Println("context tries connect timeout, terminate auth service immediately with graceful shutdown")
			w.sigint <- syscall.SIGTERM
		}
	}
}

func (w *Worker) reconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), w.reconnectTimeout)
	defer cancel()

	seconds := 100
	duration := time.Second * time.Duration(seconds)

	for {
		select {
		case <-ctx.Done():
			return ErrNoConnection
		case <-time.After(duration):
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
