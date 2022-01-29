package account

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
)

func (r *repository) GetBalance(ctx context.Context, id id.External) (currency.Currency, error) {
	const operation = "Repositories.Account.GetBalance"
	const query string = "select balance from accounts where external_id = $1"
	ret := r.db.QueryRow(ctx, query, id)
	var balance currency.Currency
	if err := ret.Scan(&balance); err != nil {
		r.logger.Error(ctx, operation, err.Error())
		if errors.Is(pgx.ErrNoRows, err) {
			return 0, account.ErrNotFound
		}
		return 0, account.ErrFetching
	}
	r.logger.Trace(ctx, operation, "finished process successfully")
	return balance, nil
}
