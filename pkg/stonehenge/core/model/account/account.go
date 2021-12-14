package account

import (
	"stonehenge/pkg/stonehenge/core/types/audits"
	"stonehenge/pkg/stonehenge/core/types/currency"
	"stonehenge/pkg/stonehenge/core/types/document"
	"stonehenge/pkg/stonehenge/core/types/id"
	"stonehenge/pkg/stonehenge/core/types/password"
)

// Account holds useful information about accounts
type Account struct {

	// The unique identifier of this entity
	Id id.ID `json:"id"`

	// The unique document that represents the user
	Document document.Document `json:"cpf"`

	// A password defined by the owner
	Secret password.Password `json:"secret"`

	// The name of the account's owner
	Name string `json:"name"`

	// The actual balance of this account
	Balance currency.Currency `json:"balance"`

	audits.Audit
}

// ToMap returns a map of this account instance
func (a *Account) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         a.Id,
		"name":       a.Name,
		"cpf":        a.Document,
		"secret":     a.Secret,
		"balance":    a.Balance,
		"updated_at": a.UpdatedAt,
		"created_at": a.CreatedAt,
	}
}

// ToArray returns an array of this account instance
func (a *Account) ToArray() []interface{} {
	return []interface{}{
		a.Id,
		a.Name,
		a.Document,
		a.Secret,
		a.Balance,
		a.UpdatedAt,
		a.CreatedAt,
	}
}
