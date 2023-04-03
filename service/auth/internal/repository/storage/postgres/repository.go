package postgres

import (
	"github.com/labstack/gommon/log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/pressly/goose"
)

func init() {
	viper.SetDefault("MIGRATIONS_DIR", "./service/auth/internal/repository/storage/postgres/migrations")
}

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

func New(db *pgxpool.Pool, o Options) (*Repository, error) {
	if err := migrations(db); err != nil {
		return nil, err
	}

	var r = &Repository{
		genSQL: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		db:     db,
	}

	r.SetOptions(o)

	return r, nil
}

func (r *Repository) SetOptions(options Options) {
	if options.DefaultLimit == 0 {
		options.DefaultLimit = 10
	}

	if options.Timeout == 0 {
		options.Timeout = time.Second * 30
	}

	if r.options != options {
		r.options = options
	}
}

func migrations(pool *pgxpool.Pool) (err error) {
	db, err := goose.OpenDBWithDriver("postgres", pool.Config().ConnConfig.ConnString())
	if err != nil {
		log.Error(err)
		return err
	}
	defer func() {
		if errClose := db.Close(); errClose != nil {
			log.Error(errClose)
			err = errClose
			return
		}
	}()

	dir := viper.GetString("MIGRATIONS_DIR")
	goose.SetTableName("user_version")
	if err = goose.Run("up", db, dir); err != nil {
		log.Error(err, zap.String("command", "up"))
		return err
	}
	return
}
