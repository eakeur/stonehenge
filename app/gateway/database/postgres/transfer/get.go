package transfer

import (
	"context"
	"stonehenge/app/core/model/transfer"
	"stonehenge/app/core/types/id"
)

func (r *repository) Get(ctx context.Context, id id.ExternalID) (transfer.Transfer, error) {
	const query string = "select * from transfers where id = $1"
	ret := r.db.QueryRow(ctx, query, id)
	tr := transfer.Transfer{}
	tr, err := parse(ret, tr)
	if err != nil {
		return tr, transfer.ErrNotFound
	}
	return tr, nil
}
