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

func CreateDatabaseURL(user, password, host, port, name, SSLMode string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, name, SSLMode)
}

// Migrate writes up the schema to a database
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

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "error uploading migration")
	}

	return nil
}

// NewConnection creates a connection object
func NewConnection(ctx context.Context, url string, _ pgx.Logger, _ pgx.LogLevel) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.ConnectConfig(ctx, pgxConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
