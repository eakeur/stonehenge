package transfer

import (
	"context"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/id"
)

func (r *repository) Get(ctx context.Context, id id.ExternalID) (transfer.Transfer, error) {
	const query string = `select
		id
		external_id
		origin_id
		destination_id
		amount
		effective_date
		updated_at
		created_at
	from transfers where id = $1`

	ret := r.db.QueryRow(ctx, query, id)
	tr := transfer.Transfer{}
	tr, err := parse(ret, tr)
	if err != nil {
		return tr, transfer.ErrNotFound
	}
	return tr, nil
}
