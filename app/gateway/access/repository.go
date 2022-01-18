package access

import (
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
}

func NewRepository(tokenExpirationTime time.Duration, tokenSigningKey []byte) access.Repository {
	return Repository{
		tokenExpirationTime: tokenExpirationTime,
		tokenSigningKey:     tokenSigningKey,
	}
}
