package persistence

import (
	"database/sql"
	"fmt"
	"stonehenge/core/repositories"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Workspace struct {
	Accounts  repositories.AccountRepository
	Transfers repositories.TransferRepository
	Identity  repositories.IdentityRepository
	db        *sql.DB
}

func (s *Workspace) Close() error {
	return s.db.Close()
}

func NewWorkspace(host, user, pass, dbname string) (*Workspace, error) {
	url := fmt.Sprintf("%v:%v@tcp(%v)/%v?parseTime=true", user, pass, host, dbname)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &Workspace{
		db: db,
		Accounts: &AccountRepository{
			*db,
		},
		Transfers: &TransferRepository{
			*db,
		},
		Identity: &IdentityRepository{
			*db,
		},
	}, nil
}
