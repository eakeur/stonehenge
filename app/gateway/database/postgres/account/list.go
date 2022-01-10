package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) List(ctx context.Context, filter account.Filter) ([]account.Account, error) {
	query := `select 
		id, 
		external_id, 
		name, 
		document, 
		balance, 
		secret, 
		updated_at, 
		created_at 
	from 
		accounts`
	args := make([]interface{}, 0)
	if filter.Name != "" {
		query = common.AppendCondition(query, "and", "name like ?")
		args = append(args, "%"+filter.Name+"%")
	}

	ret, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, account.ErrNotFound
	}
	defer ret.Close()
	accounts := make([]account.Account, 0)

	for ret.Next() {
		acc := account.Account{}
		acc, err := parse(ret, acc)
		if err != nil {
			continue
		}
		accounts = append(accounts, acc)
	}
	return accounts, nil
}
