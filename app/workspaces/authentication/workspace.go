package authentication

import (
	"context"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
)

type Workspace interface {
	// Authenticate verifies a user credential and returns the account id if it's all ok
	Authenticate(ctx context.Context, req AuthenticationRequest) (access.Access, error)
}

type workspace struct {
	ac     account.Repository
	tk     access.Manager
}

func New(ac account.Repository, tk access.Manager) *workspace {
	return &workspace{
		ac: ac,
		tk: tk,
	}
}
