package domain

import (
	model "stonehenge/model"
	"stonehenge/providers"
	"stonehenge/shared"
)

// Authenticates the user requesting access to the API. Verifies if the password is correct
// and returns a token
func Authenticate(login model.Login) (*string, error) {
	accounts, err := providers.AccountsRepository.GetAccounts(&login.Cpf)
	if err != nil {
		return nil, err
	}

	if len(accounts) == 1 {
		acc := accounts[0]
		if acc.Secret != HashSecret(login.Secret) {
			return nil, model.ErrWrongPassword
		}
		token, err := shared.CreateToken(acc.Id)
		if err != nil {
			return nil, model.ErrUnauthorized
		}
		return &token, nil
	}

	return nil, model.ErrLoginInvalid
}
