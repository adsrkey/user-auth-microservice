package transactions

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func Finish(ctx context.Context, transaction pgx.Tx, err error) error {
	if err != nil {
		if rollbackErr := transaction.Rollback(ctx); rollbackErr != nil {
			return rollbackErr
		}

		return err
	}

	commitErr := transaction.Commit(ctx)
	if commitErr != nil {
		return commitErr
	}

	return nil
}
