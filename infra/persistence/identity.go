package persistence

import (
	"database/sql"
	m "stonehenge/core/model"
	"stonehenge/infra/security"
)

type IdentityRepository struct {
	db sql.DB
}

// Authenticates the user requesting access to the API. Verifies if the password is correct
// and returns a token
func (r *IdentityRepository) Authenticate(login m.Identity) (*string, error) {
	if !login.IsCPFValid() {
		return nil, m.ErrCPFInvalid
	}

	row := SelectOne(&r.db, "accounts", "id, secret", map[string]interface{}{
		"cpf": login.Cpf,
	})

	if err := row.Err(); err != nil {
		return nil, err
	}

	var id, password string
	row.Scan(&id, &password)
	login.HashSecret()
	if err := login.ValidateHash(password); err != nil {
		return nil, err
	}

	token, err := security.CreateToken(id)
	if err != nil {
		return nil, m.ErrUnauthorized
	}
	return &token, nil

}

// Retrieves the account id by the authenticated token
func (r *IdentityRepository) GetTokenAccountId(token string) (*string, error) {
	tok, err := security.ExtractToken(token)
	if err != nil {
		return nil, err
	}
	return tok.AccountId, nil
}
