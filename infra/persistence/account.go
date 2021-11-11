package persistence

import (
	"database/sql"
	m "stonehenge/core/model"
	"time"

	"github.com/google/uuid"
)

type AccountRepository struct {
	db sql.DB
}

// Gets all accounts existing
func (r *AccountRepository) GetAll() ([]m.Account, error) {
	res, err := SelectMany(&r.db, "accounts", "*", nil)
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
	res := SelectOne(&r.db, "accounts", "*", map[string]interface{}{
		"id": id,
	})
	return r.parseRowToAccount(res)
}

// Gets the balance of the account with the ID specified
func (r *AccountRepository) GetBalanceById(id string) (*int64, error) {
	res := SelectOne(&r.db, "accounts", "balance", map[string]interface{}{
		"id": id,
	})
	if err := res.Err(); err != nil {
		return nil, err
	}
	var balance int64
	res.Scan(&balance)
	return &balance, nil
}

// Creates a new account
func (r *AccountRepository) Add(account *m.Account) (*string, error) {
	if !account.IsCPFValid() {
		return nil, m.ErrCPFInvalid
	}

	exists, err := r.Exists(account.Cpf)
	if err != nil {
		return nil, err
	}

	if !exists {
		account.Id = uuid.New().String()
		account.CreatedAt = time.Now()
		account.HashSecret()
		account.SetInitialBudget()
		_, err := Insert(&r.db, "accounts", account.ToMap())
		if err != nil {
			return nil, err
		}

		return &account.Id, nil
	}
	return nil, m.ErrAccountExists
}

// Checks if a CPF is registered to any account
func (r *AccountRepository) Exists(cpf string) (bool, error) {
	res := SelectOne(&r.db, "accounts", "count(*) as count", map[string]interface{}{
		"cpf": cpf,
	})
	var quantity int
	err := res.Scan(&quantity)
	if err != nil {
		return false, err
	}
	return quantity > 0, nil
}

// Gets the account with the ID specified
func (r *AccountRepository) Update(id string, account *m.Account) error {
	_, err := Update(&r.db, "accounts", map[string]interface{}{
		"name":    account.Name,
		"cpf":     account.Cpf,
		"secret":  account.Secret,
		"balance": account.Balance,
	}, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return err
	}

	return nil
}

// Removes an account from the
func (r *AccountRepository) Remove(id string) error {
	_, err := Delete(&r.db, "accounts", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) parseRowToAccount(row Scanner) (*m.Account, error) {
	acc := new(m.Account)
	err := row.Scan(&acc.Id, &acc.Name, &acc.Cpf, &acc.Balance, &acc.Secret, &acc.CreatedAt)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
