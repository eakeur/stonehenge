package authentication

import (
	"context"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/types/erring"
)

func (u *workspace) Authenticate(ctx context.Context, req AuthenticationRequest) (access.Access, error) {
	const operation = "Workspace.Authentication.Authentication"

	if err := req.Document.Validate(); err != nil {
		return access.Access{}, erring.Wrap(err, operation)
	}

	acc, err := u.accounts.GetWithCPF(ctx, req.Document)
	if err != nil {
		return access.Access{}, erring.Wrap(err, operation)
	}

	if err := acc.Secret.Compare(req.Secret); err != nil {
		return access.Access{}, erring.Wrap(err, operation)
	}

	tok, err := u.access.Create(acc.ExternalID)
	if err != nil {
		return access.Access{}, erring.Wrap(err, operation)
	}

	return tok, nil

}
