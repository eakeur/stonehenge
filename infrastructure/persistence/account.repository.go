package persistence

import (
	"database/sql"
	m "stonehenge/core/model"
)

type AccountRepository struct {
	db sql.DB
}

// Gets all accounts existing
func (r *AccountRepository) GetAll() ([]m.Account, error) {
	res, err := r.db.Query(MountSelect("accounts", "*", nil))
	if err != nil {
		return nil, err
	}

	defer res.Close()
	accounts := make([]m.Account, 0)

	for res.Next() {
		acc, err := r.parseRowToAccount(res)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, *acc)
	}

	return accounts, nil
}

// Gets the account with the ID specified
func (r *AccountRepository) GetById(id string) (*m.Account, error) {
	res := r.db.QueryRow(MountSelect("accounts", "*", map[string]interface{}{
		"id": id,
	}))
	return r.parseRowToAccount(res)
}

// Creates a new account
func (r *AccountRepository) Add(account *m.Account) (*string, error) {
	_, err := r.db.Exec(MountInsert("accounts", account.ToMap()))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Checks if a CPF is registered to any account
func (r *AccountRepository) Exists(cpf string) (bool, error) {
	res := r.db.QueryRow(MountSelect("accounts", "count(*) as count", map[string]interface{}{
		"cpf": cpf,
	}))
	var quantity int
	err := res.Scan(&quantity)
	if err != nil {
		return false, err
	}
	return quantity > 0, nil
}

// Gets the account with the ID specified
func (r *AccountRepository) Update(id string, account *m.Account) error {
	_, err := r.db.Exec(MountUpdate("accounts", map[string]interface{}{
		"name":    account.Name,
		"cpf":     account.Cpf,
		"secret":  account.Secret,
		"balance": account.Balance,
	}, account.ToMap()))

	if err != nil {
		return err
	}

	return nil
}

// Removes an account from the
func (r *AccountRepository) Remove(id string) error {
	_, err := r.db.Exec(MountDelete("accounts", map[string]interface{}{
		"id": id,
	}))
	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) parseRowToAccount(row Scanner) (*m.Account, error) {
	acc := new(m.Account)
	err := row.Scan(&acc.Id, &acc.Cpf, &acc.Name, &acc.Secret, &acc.Balance, &acc.CreatedAt)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
