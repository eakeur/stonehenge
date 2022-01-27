package app

import (
	"context"
	"stonehenge/app/config"
	"stonehenge/app/core/entities/access"
	contracts "stonehenge/app/core/types/logger"
	accessimpl "stonehenge/app/gateway/access"
	"stonehenge/app/gateway/database/postgres"
	"stonehenge/app/gateway/database/postgres/account"
	"stonehenge/app/gateway/database/postgres/transaction"
	"stonehenge/app/gateway/database/postgres/transfer"
	"stonehenge/app/gateway/logger"
	accountworkspace "stonehenge/app/workspaces/account"
	"stonehenge/app/workspaces/authentication"
	transferworkspace "stonehenge/app/workspaces/transfer"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Application struct {
	Accounts       accountworkspace.Workspace
	Transfers      transferworkspace.Workspace
	Authentication authentication.Workspace
	AccessManager  access.Manager
	Logger 		   contracts.Logger
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

	log := logger.NewLogger(cfg.Logger)

	accountsRepository := account.NewRepository(p, log)
	transfersRepository := transfer.NewRepository(p, log)
	transactionManager := transaction.NewManager(p)

	accountsWorkspace := accountworkspace.New(accountsRepository, transactionManager, am, log)
	transferWorkspace := transferworkspace.New(accountsRepository, transfersRepository, transactionManager, am, log)
	authenticationWorkspace := authentication.New(accountsRepository, am, log)

	return &Application{
		Accounts: accountsWorkspace, Transfers: transferWorkspace,
		AccessManager: am, Authentication: authenticationWorkspace,
		Logger: log,
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
