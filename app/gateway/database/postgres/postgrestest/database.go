package postgrestest

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/gateway/database/postgres"
)

func connect(user, password, host, port, database string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database)

	db, err := postgres.NewConnection(ctx, url, "", nil, 0)
	if err != nil {
		return db, err
	}

	err = postgres.Migrate(url, "/home/igor/go/src/stonehenge/app/gateway/database/postgres/migrations")

	return db, err
}

func purge(ctx context.Context) error {
	if err := db.Ping(ctx); err != nil {
		return err
	}

	const script = `truncate table accounts restart identity cascade;`

	_, err := db.Exec(ctx, script)
	if err != nil {
		return err
	}

	return nil
}

func RecycleDatabase(ctx context.Context) error {
	if db != nil {
		err := purge(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func PopulateAccounts(ctx context.Context, accounts ...account.Account) ([]account.Account, error) {
	const script string = `
		insert into
			accounts (document, secret, name, balance)
		values 
			($1, $2, $3, $4)
		returning 
			id, external_id, created_at, updated_at
	`

	for i, acc := range accounts {
		row := db.QueryRow(ctx, script, acc.Document, acc.Secret, acc.Name, acc.Balance)
		err := row.Scan(
			&acc.ID,
			&acc.ExternalID,
			&acc.CreatedAt,
			&acc.UpdatedAt,
		)
		if err != nil {
			return accounts, err
		}

		accounts[i] = acc
	}

	return accounts, nil
}

func PopulateTransfers(ctx context.Context, transfers ...transfer.Transfer) ([]transfer.Transfer, error) {
	const script string = `
		insert into
			transfers (account_origin_id, account_destination_id, amount, effective_date)
		values 
			($1, $2, $3, $4)
		returning 
			id, external_id, created_at, updated_at
	`
	for i, tr := range transfers {
		row := db.QueryRow(ctx, script, tr.OriginID, tr.DestinationID, tr.Amount, tr.EffectiveDate)
		err := row.Scan(
			&tr.ID,
			&tr.ExternalID,
			&tr.CreatedAt,
			&tr.UpdatedAt,
		)
		if err != nil {
			return transfers, err
		}

		transfers[i] = tr
	}

	return []transfer.Transfer{}, nil
}
