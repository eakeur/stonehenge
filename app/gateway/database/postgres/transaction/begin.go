package transaction

import (
	"context"
	"stonehenge/app/gateway/database/postgres/common"
)

func (t *pgxTransaction) Begin(ctx context.Context) (context.Context, error) {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return nil, ErrBeginTransaction
	}
	return context.WithValue(ctx, common.TXContextKey, tx), nil
}
