package transfer

import (
	"context"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) List(ctx context.Context, filter transfer.Filter) ([]transfer.Transfer, error) {
	query := "select * from transfers"
	args := make([]interface{}, 0)
	if filter.OriginID != id.ZeroValue {
		query = common.AppendCondition(query, "and", "account_origin_id = ?")
		args = append(args, filter.OriginID)
	}

	if filter.DestinationID != id.ZeroValue {
		query = common.AppendCondition(query, "and", "account_destination_id = ?")
		args = append(args, filter.DestinationID)
	}

	if !filter.InitialDate.IsZero() && !filter.FinalDate.IsZero() {
		query = common.AppendCondition(query, "and", "effective_date between ? and ?")
		args = append(args, filter.InitialDate, filter.FinalDate)

	}

	ret, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, transfer.ErrNotFound
	}
	defer ret.Close()
	transfers := make([]transfer.Transfer, 0)

	for ret.Next() {
		tr := transfer.Transfer{}
		tr, err := parse(ret, tr)
		if err != nil {
			continue
		}
		transfers = append(transfers, tr)
	}
	return transfers, nil
}
