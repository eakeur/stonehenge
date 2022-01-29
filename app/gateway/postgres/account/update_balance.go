package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/erring"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/postgres/common"
)

func (r *repository) UpdateBalance(ctx context.Context, id id.External, balance currency.Currency) error {
	const operation = "Repositories.Account.UpdateBalance"
	db, err := common.TransactionFrom(ctx)
	if err != nil {
		return erring.Wrap(err, operation)
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
		return erring.Wrap(err, operation)
	}

	if res.RowsAffected() != 1 {
		return erring.Wrap(account.ErrNotFound, operation)
	}
	return nil
}
