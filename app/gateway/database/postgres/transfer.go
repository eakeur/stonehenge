package postgres

import (
	"context"
	"stonehenge/app/core/model/transfer"
	"stonehenge/app/core/types/id"

	"github.com/jackc/pgx/v4/pgxpool"
)

type transferRepo struct {
	tx Transaction
	db *pgxpool.Pool
}

func (t *transferRepo) List(ctx context.Context, filter transfer.Filter) ([]transfer.Transfer, error) {
	query := "select * from transfers"
	args := make([]interface{}, 0)
	if filter.OriginId != "" {
		query = AppendCondition(query, "and", "account_origin_id = ?")
		args = append(args, filter.OriginId)
	}

	if filter.DestinationId != "" {
		query = AppendCondition(query, "and", "account_destination_id = ?")
		args = append(args, filter.DestinationId)
	}

	if !filter.InitialDate.IsZero() && !filter.FinalDate.IsZero() {
		query = AppendCondition(query, "and", "effective_date between ? and ?")
		args = append(args, filter.InitialDate, filter.FinalDate)

	}

	ret, err := t.db.Query(ctx, query, args...)
	if err != nil {
		return nil, transfer.ErrNotFound
	}
	defer ret.Close()
	transfers := make([]transfer.Transfer, 0)

	for ret.Next() {
		tr := transfer.Transfer{}
		tr, err := parseTransfer(ret, tr)
		if err != nil {
			continue
		}
		transfers = append(transfers, tr)
	}
	return transfers, nil
}

func (t *transferRepo) Get(ctx context.Context, id id.ExternalID) (transfer.Transfer, error) {
	const query string = "select * from transfers where id = $1"
	ret := t.db.QueryRow(ctx, query, id)
	tr := transfer.Transfer{}
	tr, err := parseTransfer(ret, tr)
	if err != nil {
		return tr, transfer.ErrNotFound
	}
	return tr, nil
}

func (t *transferRepo) Create(ctx context.Context, tran *transfer.Transfer) (id.ExternalID, error) {
	db, found := t.tx.From(ctx)
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

func parseTransfer(row Scanner, tr transfer.Transfer) (transfer.Transfer, error) {
	err := row.Scan(&tr.ID, &tr.OriginID, &tr.DestinationID, &tr.Amount, &tr.EffectiveDate, &tr.UpdatedAt, &tr.CreatedAt)
	if err != nil {
		return tr, err
	}
	return tr, nil
}

func NewTransferRepo(db *pgxpool.Pool, tx Transaction) transfer.Repository {
	return &transferRepo{
		tx, db,
	}
}
