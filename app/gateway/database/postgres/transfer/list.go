package transfer

import (
	"context"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) List(ctx context.Context, filter transfer.Filter) ([]transfer.Transfer, error) {
	query := `select
		id,
		external_id,
		origin_id,
		destination_id,
		ori.external_id,
		des.external_id,
		amount,
		effective_date,
		updated_at,
		created_at
	from 
		transfers
	inner join 
		accounts ori on origin_id = ori.id
	inner join 
		accounts des on destination_id = des.id
`
	args := make([]interface{}, 0)
	if filter.OriginID != id.ZeroValue {
		query = common.AppendCondition(query, "and", "ori.external_id = ?")
		args = append(args, filter.OriginID)
	}

	if filter.DestinationID != id.ZeroValue {
		query = common.AppendCondition(query, "and", "des.external_id = ?")
		args = append(args, filter.DestinationID)
	}

	if !filter.InitialDate.IsZero() && !filter.FinalDate.IsZero() {
		query = common.AppendCondition(query, "and", "effective_date between ? and ?")
		args = append(args, filter.InitialDate, filter.FinalDate)
	}

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
