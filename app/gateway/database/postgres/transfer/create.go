package transfer

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) Create(ctx context.Context, tr transfer.Transfer) (transfer.Transfer, error) {
	db, err := common.TransactionFrom(ctx)
	if err != nil {
		return transfer.Transfer{}, err
	}
	const script string = `
		insert into
			transfers (account_origin_id, account_destination_id, amount, effective_date)
		values 
			($1, $2, $3, $4)
		returning 
			id, external_id, created_at, updated_at
	`
	row := db.QueryRow(ctx, script, tr.OriginID, tr.DestinationID, tr.Amount, tr.EffectiveDate)
	err = row.Scan(
		&tr.ID,
		&tr.ExternalID,
		&tr.CreatedAt,
		&tr.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" && pgErr.ConstraintName == "transfers_account_origin_id_fkey" {
				return transfer.Transfer{}, transfer.ErrNonexistentOrigin
			}

			if pgErr.Code == "23503" && pgErr.ConstraintName == "transfers_account_destination_id_fkey" {
				return transfer.Transfer{}, transfer.ErrNonexistentDestination
			}
		}
		return transfer.Transfer{}, transfer.ErrRegistering
	}

	return tr, nil
}
