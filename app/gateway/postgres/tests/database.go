package tests

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"stonehenge/app/config"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/gateway/postgres"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewTestDatabase(t *testing.T) Database {
	test := t.Name()
	name := createRandomName()

	if (env == config.DatabaseConfigurations{}) {
		log.Fatal("SetupTest should be called first")
	}

	err := createDatabase(name)
	if err != nil {
		log.Fatalf("Could not setup database for test %s: %v", test, err)
	}

	pool, err := connect(name, env)
	if err != nil {
		log.Fatalf("Could not setup database for test %s: %v", test, err)
	}

	return Database{
		Name:     name,
		TestName: test,
		Pool:     pool,
		Context: context.Background(),
	}
}

type Database struct {
	Name     string
	TestName string
	Pool     *pgxpool.Pool
	Context  context.Context
}

func (d Database) Drop() {
	const script = `drop database if exists `
	d.Pool.Close()

	_, err := db.Exec(d.Context, script + d.Name)
	if err != nil {
		log.Printf("Could not drop database for test %s: %v", d.TestName, err)
	}
}

func (d Database) PopulateAccounts(ctx context.Context, accounts ...account.Account) ([]account.Account, error) {
	const script string = `
		insert into
			accounts (document, secret, name, balance)
		values 
			($1, $2, $3, $4)
		returning 
			id, external_id, created_at, updated_at
	`

	for i, acc := range accounts {
		row := d.Pool.QueryRow(ctx, script, acc.Document, acc.Secret, acc.Name, acc.Balance)
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

func (d Database) PopulateTransfers(ctx context.Context, transfers ...transfer.Transfer) ([]transfer.Transfer, error) {
	const script string = `
		insert into
			transfers (account_origin_id, account_destination_id, amount, effective_date)
		values 
			($1, $2, $3, $4)
		returning 
			id, external_id, created_at, updated_at
	`
	for i, tr := range transfers {
		row := d.Pool.QueryRow(ctx, script, tr.OriginID, tr.DestinationID, tr.Amount, tr.EffectiveDate)
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

	return transfers, nil
}

func createRandomName() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	return fmt.Sprintf("database_%d", n)
}

func createDatabase(name string) error {
	_, err := db.Exec(context.Background(), "create database " + name)
	return err
}

func connect(database string, env config.DatabaseConfigurations) (*pgxpool.Pool, error) {
	ctx := context.Background()
	url := postgres.CreateDatabaseURL(env.User, env.Password, env.Host, port, database, env.SSLMode)

	pool, err := postgres.NewConnection(ctx, url, nil, 0)
	if err != nil {
		return db, err
	}

	err = postgres.Migrate(url, env.MigrationsPath)
	return pool, err
}
