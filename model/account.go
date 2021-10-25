package model

import (
	"time"
)

// Account holds useful information about accounts
type Account struct {
	// The unique identifier of this entity
	Id string `json:"id"`

	// The name of the account's owner
	Name string `json:"name"`

	// The unique document that represents the owner of this account
	Cpf string `json:"cpf"`

	// A hash of the password defined by the owner
	Secret string `json:"secret"`

	// The actual balance of this account
	Balance int64 `json:"balance"`

	// The time the account has been created
	CreatedAt time.Time `json:"created_at"`
}

// Verifies if this account has any quantity of money
func (a *Account) HasBudget() bool {
	return a.Balance > 0.00
}

func (acc *Account) RemoveSensitiveInformation() {
	acc.Cpf = ""
	acc.Secret = ""
	acc.Balance = 0
}

// Withdraws a certain quantity from the balance of this account
func (a *Account) Withdraw(quantity int64) (int64, error) {
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

// Deposits a certain amount of money in this account
func (a *Account) Deposit(quantity int64) int64 {
	a.Balance += quantity
	return a.Balance
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

// Returns an instance of an account based on the data passed as parameter
func AccountFromMap(mapInput map[string]interface{}) Account {
	return Account{
		Id:        mapInput["id"].(string),
		Name:      mapInput["name"].(string),
		Cpf:       mapInput["cpf"].(string),
		Secret:    mapInput["secret"].(string),
		Balance:   mapInput["balance"].(int64),
		CreatedAt: mapInput["created_at"].(time.Time),
	}
}
