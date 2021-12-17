package postgres

import (
	"context"
	"github.com/google/uuid"
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
		// TODO implement error
		return nil, err
	}
	defer ret.Close()
	transfers := make([]transfer.Transfer, 0)

	for ret.Next() {
		tr, err := parseTransfer(ret)
		if err != nil {
			// TODO implement error
			return nil, err
		}
		transfers = append(transfers, *tr)
	}
	return transfers, nil
}

func (t *transferRepo) Get(ctx context.Context, id id.ID) (*transfer.Transfer, error) {
	const query string = "select * from transfers where id = $1"
	ret := t.db.QueryRow(ctx, query, id)
	tr, err := parseTransfer(ret)
	if err != nil {
		// TODO implement error
		return nil, err
	}
	return tr, nil
}

func (t *transferRepo) Create(ctx context.Context, transfer *transfer.Transfer) (*id.ID, error) {
	db, found := t.tx.From(ctx)
	if !found {
		// TODO implement missing transactions
		return nil, nil
	}
	const script string = `
		insert into
			transfers (id, account_origin_id, account_destination_id, amount, effective_date)
		values 
			($1, $2, $3, $4, $5)
		returning 
			created_at, updated_at
	`
	transfer.Id = id.ID(uuid.New().String())
	row := db.QueryRow(ctx, script, transfer.Id, transfer.OriginId, transfer.DestinationId, transfer.Amount, transfer.EffectiveDate)
	err := row.Scan(
		&transfer.CreatedAt,
		&transfer.UpdatedAt,
	)
	if err != nil {
		//TODO implement me
		return nil, err
	}

	return &transfer.Id, nil
}

func parseTransfer(row Scanner) (*transfer.Transfer, error) {
	tr := new(transfer.Transfer)
	err := row.Scan(&tr.Id, &tr.OriginId, &tr.DestinationId, &tr.Amount, &tr.EffectiveDate, &tr.UpdatedAt, &tr.CreatedAt)
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func NewTransferRepo(db *pgxpool.Pool, tx Transaction) transfer.Repository {
	return &transferRepo{
		tx, db,
	}
}