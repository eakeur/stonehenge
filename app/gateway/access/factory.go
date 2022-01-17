package access

import (
	"stonehenge/app/core/entities/access"
	"time"
)

type key int

const (
	accessContextId key = 41
)

type Factory struct {
	tokenExpirationTime time.Duration
	tokenSigningKey     []byte
}

func NewFactory(tokenExpirationTime time.Duration, tokenSigningKey []byte) access.Factory {
	return Factory{
		tokenExpirationTime: tokenExpirationTime,
		tokenSigningKey:     tokenSigningKey,
	}
}
