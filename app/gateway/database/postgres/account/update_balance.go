package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) UpdateBalance(ctx context.Context, id id.External, balance currency.Currency) error {
	db, err := common.TransactionFrom(ctx)
	if err != nil {
		return err
	}

	const script string = `
		update
			accounts
		set
			balance = $1
		where
			external_id = $2
	`
	res, err := db.Exec(ctx, script, balance, id)
	if err != nil {
		return account.ErrUpdating
	}

	if res.RowsAffected() != 1 {
		return account.ErrNotFound
	}

	return nil
}
