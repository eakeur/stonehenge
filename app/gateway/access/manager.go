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

type Manager struct {
	tokenExpirationTime time.Duration
	tokenSigningKey     []byte
	tokenSigningMethod  jwt.SigningMethod
}

func NewManager(tokenExpirationTime time.Duration, tokenSigningKey []byte) access.Manager {
	return Manager{
		tokenExpirationTime: tokenExpirationTime,
		tokenSigningKey:     tokenSigningKey,
		tokenSigningMethod:  jwt.GetSigningMethod("HS256"),
	}
}
