package postgres

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

func initDefaultEnv() error {
	if len(os.Getenv("PGHOST")) == 0 {
		if err := os.Setenv("PGHOST", "postgres"); err != nil {
			return errors.WithStack(err)
		}
	}

	if len(os.Getenv("PGPORT")) == 0 {
		if err := os.Setenv("PGPORT", "5838"); err != nil {
			return errors.WithStack(err)
		}
	}

	if len(os.Getenv("PGDATABASE")) == 0 {
		if err := os.Setenv("PGDATABASE", "postgres"); err != nil {
			return errors.WithStack(err)
		}
	}

	if len(os.Getenv("PGUSER")) == 0 {
		if err := os.Setenv("PGUSER", "postgres"); err != nil {
			return errors.WithStack(err)
		}
	}

	if len(os.Getenv("PGPASSWORD")) == 0 {
		if err := os.Setenv("PGPASSWORD", "password"); err != nil {
			return errors.WithStack(err)
		}
	}

	if len(os.Getenv("PGSSLMODE")) == 0 {
		if err := os.Setenv("PGSSLMODE", "disable"); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

type Store struct {
	Pool *pgxpool.Pool
}

type Settings struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
	SSLMode  string
}

func (s Settings) toDSN() string {
	var args []string

	if len(s.Host) > 0 {
		args = append(args, fmt.Sprintf("host=%s", s.Host))
	}

	if s.Port > 0 {
		args = append(args, fmt.Sprintf("port=%d", s.Port))
	}

	if len(s.Database) > 0 {
		args = append(args, fmt.Sprintf("dbname=%s", s.Database))
	}

	if len(s.User) > 0 {
		args = append(args, fmt.Sprintf("user=%s", s.User))
	}

	if len(s.Password) > 0 {
		args = append(args, fmt.Sprintf("password=%s", s.Password))
	}

	if len(s.SSLMode) > 0 {
		args = append(args, fmt.Sprintf("sslmode=%s", s.SSLMode))
	}

	return strings.Join(args, " ")
}

func New(settings Settings) (*Store, error) {
	if err := initDefaultEnv(); err != nil {
		panic(err)
	}

	config, err := pgxpool.ParseConfig(settings.toDSN())
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	duration := 5
	timeout := time.Second * time.Duration(duration)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return &Store{Pool: conn}, nil
}
