package persistence

import (
	"database/sql"
	m "stonehenge/core/model"
	"time"

	"github.com/google/uuid"
)

type TransferRepository struct {
	db sql.DB
}

// Gets all transfers made to or by the specified account in the ID parameter.
// The toMe bool indicates if the id passed is the destination or the origin account.
func (r *TransferRepository) GetAll(id string, toMe bool) ([]m.Transfer, error) {
	res, err := SelectMany(&r.db, "transfers", "*", nil)
	if err != nil {
		return nil, err
	}

	defer res.Close()
	transfers := make([]m.Transfer, 0)

	for res.Next() {
		acc, err := r.parseRowToTransfer(res)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, *acc)
	}

	return transfers, nil
}

// Gets the account with the ID specified
func (r *TransferRepository) GetById(id string) (*m.Transfer, error) {
	res := SelectOne(&r.db, "transfers", "*", map[string]interface{}{
		"id": id,
	})
	return r.parseRowToTransfer(res)
}

// Creates a new transfer
func (r *TransferRepository) Add(transfer *m.Transfer) (*string, error) {
	transfer.Id = uuid.New().String()
	transfer.CreatedAt = time.Now()
	_, err := Insert(&r.db, "transfers", transfer.ToMap())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Creates a money transaction between two accounts and creates a new Transfer entity
func (r *TransferRepository) Transact(transfer *m.Transfer) (*string, error) {
	transfer.Id = uuid.New().String()
	transfer.CreatedAt = time.Now()
	origin, dest, err := r.getOriginAndDestinationAccount(transfer.AccountOriginId, transfer.AccountDestinationId)
	if err != nil {
		return nil, nil
	}

	transfer.TransferMoney(origin, dest)
	tx, err := r.db.Begin()
	if err != nil {
		return nil, nil
	}

	_, err = Update(tx, "accounts", map[string]interface{}{"balance": origin.Balance}, map[string]interface{}{"id": origin.Id})
	if err != nil {
		tx.Rollback()
		return nil, nil
	}

	_, err = Update(tx, "accounts", map[string]interface{}{"balance": dest.Balance}, map[string]interface{}{"id": dest.Id})
	if err != nil {
		tx.Rollback()
		return nil, nil
	}

	_, err = Insert(tx, "transfers", transfer.ToMap())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &transfer.Id, nil
}

func (r *TransferRepository) getOriginAndDestinationAccount(originId, destId string) (*m.Account, *m.Account, error) {
	ori := new(m.Account)
	res := SelectOne(&r.db, "accounts", "id, balance", map[string]interface{}{"id": originId})
	err := res.Scan(&ori.Id, &ori.Balance)
	if err != nil {
		return nil, nil, err
	}

	dest := new(m.Account)
	res = SelectOne(&r.db, "accounts", "id, balance", map[string]interface{}{"id": destId})
	err = res.Scan(&dest.Id, &dest.Balance)
	if err != nil {
		return nil, nil, err
	}

	return ori, dest, nil
}

func (r *TransferRepository) parseRowToTransfer(row Scanner) (*m.Transfer, error) {
	tr := new(m.Transfer)
	err := row.Scan(&tr.Id, &tr.AccountOriginId, &tr.AccountDestinationId, &tr.Amount, &tr.CreatedAt)
	if err != nil {
		return nil, err
	}
	return tr, nil
}
