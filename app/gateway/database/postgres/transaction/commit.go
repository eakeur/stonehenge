package transaction

import (
	"context"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/gateway/database/postgres/common"
)

func (t *pgxTransaction) Commit(ctx context.Context) error {
	tx, ok := common.TransactionFrom(ctx)
	if !ok {
		return transaction.ErrNoTransaction
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
