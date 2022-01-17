package access

import (
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/types/id"

	"github.com/golang-jwt/jwt"
)

func (f Factory) ExtractAccessFromToken(token string) (access.Access, error) {
	parsed, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return f.tokenSigningKey, nil
	})
	if err != nil {
		return access.Access{}, access.ErrTokenFailedCreation
	}

	if parsed.Claims.Valid() != nil {
		return access.Access{}, access.ErrTokenInvalidOrExpired
	}
	claims := parsed.Claims.(jwt.StandardClaims)

	ext := id.ExternalFrom(claims.Id)

	if ext == id.Zero {
		return access.Access{}, id.ErrInvalidID
	}

	return access.Access{
		AccountID: ext,
		Token:     token,
	}, nil
}
