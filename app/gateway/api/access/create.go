package access

import (
	"github.com/golang-jwt/jwt"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
	"time"
)

func (f Factory) Create(acc account.Account) (access.Access, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))
	t.Claims = &jwt.StandardClaims{
		Id:        acc.ExternalID.String(),
		ExpiresAt: time.Now().Add(f.tokenExpirationTime).Unix(),
	}

	token, err := t.SignedString(f.tokenSigningKey)
	if err != nil {
		return access.Access{}, access.ErrTokenFailedCreation
	}

	return access.Access{
		AccountID: acc.ExternalID,
		Token:     token,
	}, nil
}
