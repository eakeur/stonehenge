package transaction

import (
	"context"
	"stonehenge/app/gateway/database/postgres/common"
)

func (t *manager) Begin(ctx context.Context) (context.Context, error) {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, common.TXContextKey, tx), nil
}
