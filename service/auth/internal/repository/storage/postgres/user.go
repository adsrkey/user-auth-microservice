package postgres

import (
	"auth-service/pkg/tool/transactions"
	"auth-service/service/auth/internal/domain/user"
	"auth-service/service/auth/internal/repository/storage/postgres/dao"
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

func (r *Repository) GetUser(ctx context.Context, user *user.User) (*dao.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.options.Timeout)
	defer cancel()

	transaction, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, transaction pgx.Tx) {
		err = transactions.Finish(ctx, transaction, err)
	}(ctx, transaction)

	response, err := r.getUserTx(ctx, transaction, user)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) getUserTx(ctx context.Context, tx pgx.Tx, user *user.User) (*dao.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.options.Timeout)
	defer cancel()

	response, err := r.oneUserTx(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) oneUserTx(ctx context.Context, transaction pgx.Tx, user *user.User) (*dao.User, error) {
	var builder = r.genSQL.Select(
		"id",
		"created_at",
		"modified_at",
		"email",
		"hash",
	).From("users.users")

	builder = builder.Where(squirrel.Eq{"email": user.Email})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := transaction.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoUser dao.User
	if err = pgxscan.ScanOne(&daoUser, rows); err != nil {
		return nil, err
	}

	return &daoUser, nil
}

func (r Repository) toCopyFromSource(contacts ...*user.User) pgx.CopyFromSource {
	rows := make([][]interface{}, len(contacts))

	for i, val := range contacts {
		rows[i] = []interface{}{
			time.Now(),
			time.Now(),
			val.Email,
			val.Hash,
		}
	}

	return pgx.CopyFromRows(rows)
}

func (r *Repository) CreateUser(ctx context.Context, user *user.User) error {
	ctx, cancel := context.WithTimeout(ctx, r.options.Timeout)
	defer cancel()

	transaction, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(ctx context.Context, transaction pgx.Tx) {
		err = transactions.Finish(ctx, transaction, err)
	}(ctx, transaction)

	err = r.createUserTx(ctx, transaction, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) createUserTx(ctx context.Context, tx pgx.Tx, user *user.User) error {
	ctx, cancel := context.WithTimeout(ctx, r.options.Timeout)
	defer cancel()

	err := r.createOneUserTx(ctx, tx, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) createOneUserTx(ctx context.Context, tx pgx.Tx, user *user.User) error {
	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"users", "users"},
		CreateColumnContact,
		r.toCopyFromSource(user))
	if err != nil {
		return err
	}

	return nil
}

var CreateColumnContact = []string{
	"created_at",
	"modified_at",
	"email",
	"hash",
}
