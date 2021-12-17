package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

// Migrate - Write up the schema to a database
func Migrate(db string, filesPath string) error {
	path := fmt.Sprintf("file://%v", filesPath)
	migration, err := migrate.New(path, db)
	if err != nil {
		return err
	}
	defer migration.Close()

	if err != nil {
		return errors.Wrap(err, "could not start migration")
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange && err != migrate.ErrInvalidVersion {
		return errors.Wrap(err, "error uploading migration")
	}

	return nil
}

// NewConnection creates a connection object and runs a migration in this connection
func NewConnection(ctx context.Context, url, migrationsPath string, log pgx.Logger, level pgx.LogLevel) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	//pgxConfig.ConnConfig.Logger = log

	//pgxConfig.ConnConfig.LogLevel = level

	db, err := pgxpool.ConnectConfig(ctx, pgxConfig)
	if err != nil {
		return nil, err
	}

	err = Migrate(url, migrationsPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}
