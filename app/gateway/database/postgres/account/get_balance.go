package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
)

func (r *repository) GetBalance(ctx context.Context, id id.ExternalID) (currency.Currency, error) {
	const query string = "select balance from accounts where external_id = $1"
	ret := r.db.QueryRow(ctx, query, id)
	var balance currency.Currency
	if err := ret.Scan(&balance); err != nil {
		return 0, account.ErrNotFound
	}
	return balance, nil
}
