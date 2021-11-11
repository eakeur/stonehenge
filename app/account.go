package app

import (
	"stonehenge/app/dto"
	m "stonehenge/core/model"
	"stonehenge/core/repositories"
)

type AccountApp interface {

	// Gets all accounts existing
	GetAll() ([]dto.SafeAccountDTO, error)

	// Gets the account with the ID specified
	GetById(id string) (*dto.AccountDTO, error)

	// Gets the balance of the account with the ID specified
	GetBalanceById(id string) (*dto.BalanceDTO, error)

	// Creates a new account
	Add(account *m.Account) (*string, error)

	// Checks if a CPF is registered to any account
	Exists(cpf string) (bool, error)

	// Gets the account with the ID specified
	Update(id string, account *m.Account) error

	// Removes an account from the
	Remove(id string) error
}

type accountApp struct {
	ac repositories.AccountRepository
}

func NewAccountApp(accountsRepo repositories.AccountRepository) AccountApp {
	return &accountApp{
		ac: accountsRepo,
	}
}

func (r *accountApp) GetAll() ([]dto.SafeAccountDTO, error) {
	res, err := r.ac.GetAll()
	if err != nil {
		return nil, err
	}
	accounts := make([]dto.SafeAccountDTO, len(res))
	for i, acc := range res {
		accounts[i] = dto.SafeAccountDTO{Id: acc.Id, Name: acc.Name}
	}
	return accounts, nil
}

// Gets the account with the ID specified
func (r *accountApp) GetById(id string) (*dto.AccountDTO, error) {
	res, err := r.ac.GetById(id)
	if err != nil {
		return nil, err
	}

	return &dto.AccountDTO{Id: res.Id, Name: res.Name, Cpf: res.Cpf, Secret: res.Secret, Balance: res.Balance}, nil
}

// Gets the balance of the account with the ID specified
func (r *accountApp) GetBalanceById(id string) (*dto.BalanceDTO, error) {
	balance, err := r.ac.GetBalanceById(id)
	if err != nil {
		return nil, err
	}
	return &dto.BalanceDTO{
		Balance: *balance,
	}, nil
}

// Creates a new account
func (r *accountApp) Add(account *m.Account) (*string, error) {
	return r.ac.Add(account)
}

// Checks if a CPF is registered to any account
func (r *accountApp) Exists(cpf string) (bool, error) {
	return r.ac.Exists(cpf)
}

// Gets the account with the ID specified
func (r *accountApp) Update(id string, account *m.Account) error {
	return r.ac.Update(id, account)
}

// Removes an account from the
func (r *accountApp) Remove(id string) error {
	return r.ac.Remove(id)
}

var _ AccountApp = &accountApp{}
