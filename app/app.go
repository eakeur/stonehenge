package app

import (
	"context"
	"stonehenge/app/config"
	"stonehenge/app/core/entities/access"
	accessimpl "stonehenge/app/gateway/access"
	"stonehenge/app/gateway/database/postgres"
	"stonehenge/app/gateway/database/postgres/account"
	"stonehenge/app/gateway/database/postgres/transaction"
	"stonehenge/app/gateway/database/postgres/transfer"
	accountworkspace "stonehenge/app/workspaces/account"
	"stonehenge/app/workspaces/authentication"
	transferworkspace "stonehenge/app/workspaces/transfer"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Application struct {
	Accounts       accountworkspace.Workspace
	Transfers      transferworkspace.Workspace
	Authentication authentication.Workspace
	AccessManager  access.Manager
}

func NewApplication(ctx context.Context, cfg config.Config) (*Application, error) {

	p, err := setupDB(ctx, cfg.Database)
	if err != nil {
		return &Application{}, err
	}

	am, err := setupAccessManager(cfg.Access)
	if err != nil {
		return &Application{}, err
	}

	accountsRepository := account.NewRepository(p)
	transfersRepository := transfer.NewRepository(p)
	transactionManager := transaction.NewManager(p)

	accountsWorkspace := accountworkspace.New(accountsRepository, transactionManager, am)
	transferWorkspace := transferworkspace.New(accountsRepository, transfersRepository, transactionManager, am)
	authenticationWorkspace := authentication.New(accountsRepository, am)

	return &Application{
		Accounts: accountsWorkspace, Transfers: transferWorkspace,
		AccessManager: am, Authentication: authenticationWorkspace,
	}, nil

}

func setupAccessManager(cfg config.AccessConfigurations) (access.Manager, error) {
	duration, err := strconv.Atoi(cfg.ExpirationTime)
	if err != nil {
		return accessimpl.Manager{}, errors.Wrap(err, "failed parsing access expiration time")
	}

	return accessimpl.NewManager(time.Minute*time.Duration(duration), []byte(cfg.SigningKey)), err
}

func setupDB(ctx context.Context, cfg config.DatabaseConfigurations) (*pgxpool.Pool, error) {

	url := postgres.CreateDatabaseURL(cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)
	conn, err := postgres.NewConnection(ctx, url, nil, 5)
	if err != nil {
		return conn, err
	}

	err = postgres.Migrate(url, cfg.MigrationsPath)
	if err != nil {
		conn.Close()
		return conn, err
	}

	return conn, nil
}
