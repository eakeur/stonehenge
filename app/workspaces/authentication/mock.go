package authentication

import (
	"context"
	"stonehenge/app/core/entities/access"
)

var _ Workspace = workspaceMock{}

type workspaceMock struct {
	AuthenticateResult access.Access
	Error error
}

func (w workspaceMock) Authenticate(_ context.Context, _ AuthenticationRequest) (access.Access, error) {
	return w.AuthenticateResult, w.Error
}

