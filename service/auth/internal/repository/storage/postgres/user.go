package postgres

import (
	"auth-service/pkg/tool/transaction"
	context "auth-service/pkg/type"
	"auth-service/service/auth/internal/domain/user"
	"auth-service/service/auth/internal/repository/storage/postgres/dao"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"time"
)

func (r *Repository) GetUser(c context.Context, user *user.User) (*dao.User, error) {
	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	response, err := r.getUserTx(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) getUserTx(c context.Context, tx pgx.Tx, user *user.User) (*dao.User, error) {
	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	response, err := r.oneUserTx(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) oneUserTx(ctx context.Context, tx pgx.Tx, user *user.User) (*dao.User, error) {
	var builder = r.genSQL.Select(
		"id",
		"created_at",
		"modified_at",
		"email",
		"hash",
	).From("adsrkey.user")

	builder = builder.Where(squirrel.Eq{"email": user.Email})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, query, args...)
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

func (r *Repository) toDomainUser(in *dao.User) (*user.User, error) {
	// TODO
	u := &user.User{
		Email: in.Email,
		Hash:  in.Hash,
	}
	return u, nil
}

func (r *Repository) CreateUser(c context.Context, user *user.User) error {
	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	err = r.createUserTx(ctx, tx, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) createUserTx(c context.Context, tx pgx.Tx, user *user.User) error {
	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	err = r.createOneUserTx(ctx, tx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) createOneUserTx(ctx context.Context, tx pgx.Tx, user *user.User) error {
	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"adsrkey", "user"},
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
