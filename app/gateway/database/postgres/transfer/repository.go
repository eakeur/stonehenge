package transfer

import (
	"stonehenge/app/core/model/transfer"
	"stonehenge/app/gateway/database/postgres/common"
	"stonehenge/app/gateway/database/postgres/transaction"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	tx transaction.Transaction
	db *pgxpool.Pool
}

func parse(row common.Scanner, tr transfer.Transfer) (transfer.Transfer, error) {
	err := row.Scan(&tr.ID, &tr.OriginID, &tr.DestinationID, &tr.Amount, &tr.EffectiveDate, &tr.UpdatedAt, &tr.CreatedAt)
	if err != nil {
		return tr, err
	}
	return tr, nil
}

func NewRepository(db *pgxpool.Pool, tx transaction.Transaction) transfer.Repository {
	return &repository{
		tx, db,
	}
}
