package postgres

import (
	"time"

	"github.com/labstack/gommon/log"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/pressly/goose"
)

type Repository struct {
	db      *pgxpool.Pool
	genSQL  squirrel.StatementBuilderType
	options Options
}

type Options struct {
	Timeout       time.Duration
	DefaultLimit  uint64
	DefaultOffset uint64
}

func New(pool *pgxpool.Pool, options Options) (*Repository, error) {
	viper.SetDefault("MIGRATIONS_DIR", "service/auth/internal/repository/storage/postgres/migrations")

	if err := migrations(pool); err != nil {
		return nil, err
	}

	var repository = &Repository{
		genSQL: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		db:     pool,
	}

	repository.SetOptions(options)

	return repository, nil
}

func (r *Repository) SetOptions(options Options) {
	if options.DefaultLimit == 0 {
		options.DefaultLimit = 10
	}

	if options.Timeout == 0 {
		timeout := 30
		duration := time.Second * time.Duration(timeout)
		options.Timeout = duration
	}

	if r.options != options {
		r.options = options
	}
}

func migrations(pool *pgxpool.Pool) (err error) {
	database, err := goose.OpenDBWithDriver("postgres", pool.Config().ConnConfig.ConnString())
	if err != nil {
		log.Error(err)

		return err
	}

	defer func() {
		if errClose := database.Close(); errClose != nil {
			log.Error(errClose)
			err = errClose

			return
		}
	}()

	dir := viper.GetString("MIGRATIONS_DIR")

	goose.SetTableName("user_version")

	if err = goose.Run("up", database, dir); err != nil {
		log.Error(err, zap.String("command", "up"))

		return err
	}

	return
}
