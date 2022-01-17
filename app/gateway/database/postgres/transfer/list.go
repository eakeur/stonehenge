package transfer

import (
	"context"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) List(ctx context.Context, filter transfer.Filter) ([]transfer.Transfer, error) {
	query := `select
		t.id,
		t.external_id,
		t.account_origin_id,
		t.account_destination_id,
		ori.external_id,
		des.external_id,
		t.amount,
		t.effective_date,
		t.updated_at,
		t.created_at
	from 
		transfers t
	inner join 
		accounts ori on account_origin_id = ori.id
	inner join 
		accounts des on account_destination_id = des.id
`
	args := make([]interface{}, 0)
	idx := 0
	if filter.OriginID != id.Zero {
		idx++
		query = common.AppendCondition(query, "and", "ori.external_id = ?", idx)
		args = append(args, filter.OriginID)
	}

	if filter.DestinationID != id.Zero {
		idx++
		query = common.AppendCondition(query, "and", "des.external_id = ?", idx)
		args = append(args, filter.DestinationID)
	}

	if !filter.InitialDate.IsZero() && !filter.FinalDate.IsZero() {
		idx += 2
		query = common.AppendCondition(query, "and", "effective_date between ? and ?", idx-1, idx)
		args = append(args, filter.InitialDate, filter.FinalDate)
	}

	query += "\n order by effective_date desc"

	ret, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return []transfer.Transfer{}, transfer.ErrFetching
	}
	defer ret.Close()

	transfers := make([]transfer.Transfer, 0)
	for ret.Next() {
		tr := transfer.Transfer{}
		err = ret.Scan(
			&tr.ID,
			&tr.ExternalID,
			&tr.OriginID,
			&tr.DestinationID,
			&tr.Details.OriginExternalID,
			&tr.Details.DestinationExternalID,
			&tr.Amount,
			&tr.EffectiveDate,
			&tr.UpdatedAt,
			&tr.CreatedAt)
		transfers = append(transfers, tr)
	}
	return transfers, nil
}
