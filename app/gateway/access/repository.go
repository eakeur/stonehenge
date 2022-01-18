package access

import (
	"github.com/golang-jwt/jwt"
	"stonehenge/app/core/entities/access"
	"time"
)

type key int

const (
	accessContextId key = 41
)

type Repository struct {
	tokenExpirationTime time.Duration
	tokenSigningKey     []byte
	tokenSigningMethod  jwt.SigningMethod
}

func NewRepository(tokenExpirationTime time.Duration, tokenSigningKey []byte) access.Repository {
	return Repository{
		tokenExpirationTime: tokenExpirationTime,
		tokenSigningKey:     tokenSigningKey,
		tokenSigningMethod: jwt.GetSigningMethod("HS256"),
	}
}
