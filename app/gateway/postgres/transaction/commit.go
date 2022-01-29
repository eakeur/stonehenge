package transaction

import (
	"context"
	"stonehenge/app/gateway/postgres/common"
)

func (t *manager) Commit(ctx context.Context) error {
	tx, err := common.TransactionFrom(ctx)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
