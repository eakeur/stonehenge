package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/document"
)

func (r *repository) GetWithCPF(ctx context.Context, document document.Document) (account.Account, error) {
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
	where 
		document = $1`

	ret := r.db.QueryRow(ctx, query, document)
	acc := account.Account{}
	acc, err := parse(ret, acc)
	if err != nil {
		return acc, account.ErrNotFound
	}
	return acc, nil
}
