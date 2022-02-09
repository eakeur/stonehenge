package authentication

import (
	"context"
	"stonehenge/app/core/entities/access"
)

var _ Workspace = WorkspaceMock{}

type WorkspaceMock struct {
	AuthenticateResult access.Access
	Error              error
}

func (w WorkspaceMock) Authenticate(_ context.Context, _ AuthenticationRequest) (access.Access, error) {
	return w.AuthenticateResult, w.Error
}
