package app

import (
	m "stonehenge/core/model"
	"stonehenge/core/repositories"
)

type IdentityApp interface {
	// Authenticates the user requesting access to the API. Verifies if the password is correct
	// and returns a token
	Authenticate(login m.Identity) (*string, error)

	// Retrieves the account id by the authenticated token
	GetTokenAccountId(token string) (*string, error)
}

type identityApp struct {
	lg repositories.IdentityRepository
}

func NewIdentityApp(loginRepo repositories.IdentityRepository) IdentityApp {
	return &identityApp{
		lg: loginRepo,
	}
}

// Authenticates the user requesting access to the API. Verifies if the password is correct
// and returns a token
func (r *identityApp) Authenticate(login m.Identity) (*string, error) {
	return r.lg.Authenticate(login)

}

// Retrieves the account id by the authenticated token
func (r *identityApp) GetTokenAccountId(token string) (*string, error) {
	return r.lg.GetTokenAccountId(token)
}

var _ IdentityApp = &identityApp{}
