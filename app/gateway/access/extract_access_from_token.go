package access

import (
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/types/erring"
	"stonehenge/app/core/types/id"

	"github.com/golang-jwt/jwt"
)

func (f Manager) ExtractAccessFromToken(token string) (access.Access, error) {
	const operation = "Managers.Access.ExtractAccessFromToken"
	parsed, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != f.tokenSigningMethod {
			return access.Access{}, access.ErrTokenInvalidOrExpired
		}
		return f.tokenSigningKey, nil
	})
	if err != nil {
		return access.Access{}, erring.Wrap(access.ErrTokenInvalidOrExpired, operation)
	}

	if parsed.Claims.Valid() != nil {
		return access.Access{}, erring.Wrap(access.ErrTokenInvalidOrExpired, operation)
	}
	claims := parsed.Claims.(*jwt.StandardClaims)

	ext := id.ExternalFrom(claims.Id)

	if ext == id.Zero {
		return access.Access{}, erring.Wrap(id.ErrInvalidID, operation)
	}

	return access.Access{
		AccountID: ext,
		Token:     token,
	}, nil
}
