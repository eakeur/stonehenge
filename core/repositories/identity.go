package repositories

import (
	m "stonehenge/core/model"
)

type IdentityRepository interface {
	// Authenticates the user requesting access to the API. Verifies if the password is correct
	// and returns a token
	Authenticate(login m.Identity) (*string, error)

	// Retrieves the account id by the authenticated token
	GetTokenAccountId(token string) (*string, error)
}
