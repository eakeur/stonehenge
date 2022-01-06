package transaction

import (

	"github.com/jackc/pgx/v4/pgxpool"
	"stonehenge/app/core/entities/transaction"
)

type pgxTransaction struct {
	db *pgxpool.Pool
}

// NewTransaction creates a transaction adapter object
func NewTransaction(db *pgxpool.Pool) transaction.Transaction {
	return &pgxTransaction{
		db: db,
	}
}
