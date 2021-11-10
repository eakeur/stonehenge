package repositories

import m "stonehenge/core/model"

type TransferRepository interface {

	// Gets all transfers made to or by the specified account in the ID parameter.
	// The toMe bool indicates if the id passed is the destination or the origin account.
	GetAll(id string, toMe bool) ([]m.Transfer, error)

	// Gets the account with the ID specified
	GetById(id string) (*m.Transfer, error)

	// Creates a new transfer
	Add(transfer *m.Transfer) (*string, error)

	// Creates a money transaction between two accounts and creates a new Transfer entity
	UpdateAccountsInTransaction(transfer *m.Transfer, origin *m.Account, destination *m.Account) (*string, error)
}
