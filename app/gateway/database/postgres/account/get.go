package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/id"
)

func (r *repository) Get(ctx context.Context, id id.ExternalID) (account.Account, error) {
	const query string = `select 
		id, 
		external_id, 
		name, 
		document, 
		balance, 
		secret, 
		updated_at, 
		created_at 
	from 
		accounts
	where external_id = $1`

	acc := account.Account{}
	ret := r.db.QueryRow(ctx, query, id)
	acc, err := parse(ret, acc)
	if err != nil {
		return acc, account.ErrNotFound
	}
	return acc, nil
}
