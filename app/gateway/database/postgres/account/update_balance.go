package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) UpdateBalance(ctx context.Context, id id.ExternalID, balance currency.Currency) error {
	db, found := common.TransactionFrom(ctx)
	if !found {
		return account.ErrCreating
	}

	const script string = `
		update
			accounts
		set
			balance = $1
		where
			id = $2
	`
	_, err := db.Exec(ctx, script, balance, id)
	if err != nil {
		return account.ErrCreating
	}

	return nil
}
