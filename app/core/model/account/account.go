package account

import (
	"stonehenge/app/core/types/audits"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
)

// Account holds useful information about accounts
type Account struct {

	// The unique identifier of this entity
	ID id.ID `json:"id"`

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
