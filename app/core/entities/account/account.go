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
	ID id.ID

	// ExternalID is the public identifier of this entity
	ExternalID id.External

	// The unique document that represents the user
	Document document.Document

	// A password defined by the owner
	Secret password.Password

	// The name of the account's owner
	Name string

	// The actual balance of this account
	Balance currency.Currency

	audits.Audit
}

func (a Account) Validate() error {

	// Checks document's consistency
	if err := a.Document.Validate(); err != nil {
		return err
	}

	return nil

}
