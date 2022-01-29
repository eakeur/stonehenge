package transaction

import (
	"context"
	"stonehenge/app/gateway/postgres/common"
)

func (t *manager) Begin(ctx context.Context) context.Context {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return ctx
	}
	return context.WithValue(ctx, common.TXContextKey, tx)
}
