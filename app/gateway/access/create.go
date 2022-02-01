package access

import (
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/types/id"
	"time"

	"github.com/golang-jwt/jwt"
)

func (f Manager) Create(ext id.External, name string) (access.Access, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))
	t.Claims = &jwt.StandardClaims{
		Id:        ext.String(),
		ExpiresAt: time.Now().Add(f.tokenExpirationTime).Unix(),
		Subject:   name,
	}

	token, err := t.SignedString(f.tokenSigningKey)
	if err != nil {
		return access.Access{}, access.ErrTokenFailedCreation
	}

	return access.Access{
		AccountID:   ext,
		AccountName: name,
		Token:       token,
	}, nil
}
