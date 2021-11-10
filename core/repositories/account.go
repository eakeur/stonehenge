package repositories

import m "stonehenge/core/model"

type AccountRepository interface {

	// Gets all accounts existing
	GetAll() ([]m.Account, error)

	// Gets the account with the ID specified
	GetById(id string) (*m.Account, error)

	// Creates a new account
	Add(account *m.Account) (*string, error)

	// Checks if a CPF is registered to any account
	Exists(cpf string) (bool, error)

	// Gets the account with the ID specified
	Update(id string, account *m.Account) error

	// Removes an account from the
	Remove(id string) error
}
