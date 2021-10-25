package domain

import (
	model "stonehenge/model"
	"stonehenge/providers"
	"time"

	"github.com/golang-jwt/jwt"
)

const JWT_KEY string = "EDAF12D5D997C58B1962FD8350E8B1C158447B5D1002DABA4F551BC3CD38F236"

const TOKEN_VALID_MINUTES = 10

// Authenticates the user requesting access to the API. Verifies if the password is correct
// and returns a token
func Authenticate(login model.Login) (*string, error) {
	login.Cpf = NormalizeCPF(login.Cpf)
	if !IsCPFValid(login.Cpf) {
		return nil, model.ErrCPFInvalid
	}
	accounts, err := providers.AccountsRepository.GetAccounts(&login.Cpf)
	if err != nil {
		return nil, err
	}

	if len(accounts) == 1 {
		acc := accounts[0]
		if acc.Secret != HashSecret(login.Secret) {
			return nil, model.ErrWrongPassword
		}
		token, err := CreateToken(acc.Id)
		if err != nil {
			return nil, model.ErrUnauthorized
		}
		return &token, nil
	}

	return nil, model.ErrLoginInvalid
}

// Retrieves the account id by the authenticated token
func GetAccountIdByToken(token string) (*string, error) {
	parsed, err := jwt.ParseWithClaims(token, &TokenDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_KEY), nil
	})
	if err != nil {
		return nil, model.ErrUnauthorized
	}

	if parsed.Claims.Valid() != nil {
		return nil, model.ErrUnauthorized
	}
	claims := parsed.Claims.(*TokenDetails)
	return &claims.AccountId, nil
}

// Creates a JWT token string containing the account id of the applicant
func CreateToken(userId string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))
	t.Claims = TokenDetails{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * TOKEN_VALID_MINUTES).Unix(),
		},
		userId,
	}
	return t.SignedString([]byte(JWT_KEY))
}

// A struct containing the type of data in the parsed token
type TokenDetails struct {
	*jwt.StandardClaims

	// The id of the current account
	AccountId string
}
