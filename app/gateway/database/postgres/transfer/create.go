package transfer

import (
	"context"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) Create(ctx context.Context, tran *transfer.Transfer) (id.ExternalID, error) {
	db, found := common.TransactionFrom(ctx)
	if !found {
		return id.New(), transfer.ErrRegistering
	}
	const script string = `
		insert into
			transfers (id, account_origin_id, account_destination_id, amount, effective_date)
		values 
			($1, $2, $3, $4, $5)
		returning 
			id, external_id, created_at, updated_at
	`
	row := db.QueryRow(ctx, script, tran.ID, tran.OriginID, tran.DestinationID, tran.Amount, tran.EffectiveDate)
	err := row.Scan(
		&tran.ID,
		&tran.ExternalID,
		&tran.CreatedAt,
		&tran.UpdatedAt,
	)
	if err != nil {
		return id.New(), transfer.ErrRegistering
	}

	return tran.ExternalID, nil
}
