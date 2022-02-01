package transfer

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/erring"
	"stonehenge/app/gateway/postgres/common"
)

func (r *repository) Create(ctx context.Context, tr transfer.Transfer) (transfer.Transfer, error) {
	const operation = "Repositories.Transfer.Create"
	db, err := common.TransactionFrom(ctx)
	if err != nil {
		return transfer.Transfer{}, erring.Wrap(err, operation)
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
			if pgErr.Code == common.PostgresNonexistentFK && pgErr.ConstraintName == "transfers_account_origin_id_fkey" {
				return transfer.Transfer{}, erring.Wrap(transfer.ErrNonexistentOrigin, operation)
			}

			if pgErr.Code == common.PostgresNonexistentFK && pgErr.ConstraintName == "transfers_account_destination_id_fkey" {
				return transfer.Transfer{}, erring.Wrap(transfer.ErrNonexistentDestination, operation)
			}
		}
		return transfer.Transfer{}, erring.Wrap(err, operation)
	}
	return tr, nil
}
