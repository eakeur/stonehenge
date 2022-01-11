package transfer

import (
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/gateway/database/postgres/common"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func parse(row common.Scanner, tr transfer.Transfer) (transfer.Transfer, error) {
	err := row.Scan(&tr.ID, &tr.ExternalID, &tr.OriginID, &tr.DestinationID, &tr.Amount, &tr.EffectiveDate, &tr.UpdatedAt, &tr.CreatedAt)
	if err != nil {
		return tr, err
	}
	return tr, nil
}

func NewRepository(db *pgxpool.Pool) transfer.Repository {
	return &repository{
		db: db,
	}
}
