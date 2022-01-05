package account

import (
	"context"
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
)

func (r *repository) UpdateBalance(ctx context.Context, id id.ExternalID, balance currency.Currency) error {
	db, found := r.tx.From(ctx)
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
