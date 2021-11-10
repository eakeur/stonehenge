package persistence

import (
	"database/sql"
	m "stonehenge/core/model"
)

type TransferRepository struct {
	db sql.DB
}

// Gets all transfers made to or by the specified account in the ID parameter.
// The toMe bool indicates if the id passed is the destination or the origin account.
func (r *TransferRepository) GetAll(id string, toMe bool) ([]m.Transfer, error) {
	res, err := r.db.Query(MountSelect("transfers", "*", nil))
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
	res := r.db.QueryRow(MountSelect("transfers", "*", map[string]interface{}{
		"id": id,
	}))
	return r.parseRowToTransfer(res)
}

// Creates a new transfer
func (r *TransferRepository) Add(transfer *m.Transfer) (*string, error) {
	_, err := r.db.Exec(MountInsert("transfers", transfer.ToMap()))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Creates a money transaction between two accounts and creates a new Transfer entity
func (r *TransferRepository) UpdateAccountsInTransaction(transfer *m.Transfer, origin *m.Account, destination *m.Account) (*string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, nil
	}

	_, err = tx.Exec(MountUpdate("accounts", map[string]interface{}{"balance": origin.Balance}, map[string]interface{}{"id": origin.Id}))
	if err != nil {
		tx.Rollback()
		return nil, nil
	}

	_, err = tx.Exec(MountUpdate("accounts", map[string]interface{}{"balance": destination.Balance}, map[string]interface{}{"id": destination.Id}))
	if err != nil {
		tx.Rollback()
		return nil, nil
	}

	_, err = tx.Exec(MountInsert("transfers", transfer.ToMap()))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return nil, nil
}

func (r *TransferRepository) parseRowToTransfer(row Scanner) (*m.Transfer, error) {
	tr := new(m.Transfer)
	err := row.Scan(&tr.Id, &tr.AccountOriginId, &tr.AccountDestinationId, &tr.Amount, &tr.CreatedAt)
	if err != nil {
		return nil, err
	}
	return tr, nil
}
