package model

import (
	"time"
)

// How much money will the users get when creating a new account
const INITIAL_BUDGET int = 50000

// Account holds useful information about accounts
type Account struct {
	Identity

	// The unique identifier of this entity
	Id string `json:"id"`

	// The name of the account's owner
	Name string `json:"name"`

	// The actual balance of this account
	Balance int64 `json:"balance"`

	// The time the account has been created
	CreatedAt time.Time `json:"created_at"`
}

// Verifies if this account has any quantity of money
func (a *Account) HasBudget() bool {
	return a.Balance > 0.00
}

// Withdraws a certain quantity from the balance of this account
func (a *Account) Withdraw(quantity int64) (int64, error) {
	if quantity > 0 {
		if a.HasBudget() {
			remainingMoney := a.Balance - quantity
			if remainingMoney < 0 {
				return 0, ErrNoMoney
			}
			a.Balance = remainingMoney
			return remainingMoney, nil
		}
		return 0, ErrNoMoney
	}
	return 0, ErrAmountInvalid
}

// Deposits a certain amount of money in this account
func (a *Account) Deposit(quantity int64) int64 {
	a.Balance += quantity
	return a.Balance
}

// Deposits a certain amount of money in this account when its first created
func (a *Account) SetInitialBudget() {
	a.Deposit(int64(INITIAL_BUDGET))
}

// Returns a map of this account instance
func (a *Account) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         a.Id,
		"name":       a.Name,
		"cpf":        a.Cpf,
		"secret":     a.Secret,
		"balance":    a.Balance,
		"created_at": a.CreatedAt,
	}
}

// Returns an array of this account instance
func (a *Account) ToArray() []interface{} {
	return []interface{}{
		a.Id,
		a.Name,
		a.Cpf,
		a.Secret,
		a.Balance,
		a.CreatedAt,
	}
}

// Returns an instance of an account based on the data passed as parameter
func AccountFromMap(mapInput map[string]interface{}) Account {
	return Account{

		Id:   mapInput["id"].(string),
		Name: mapInput["name"].(string),
		Identity: Identity{
			Cpf:    mapInput["cpf"].(string),
			Secret: mapInput["cpf"].(string),
		},
		Balance:   mapInput["balance"].(int64),
		CreatedAt: mapInput["created_at"].(time.Time),
	}
}
