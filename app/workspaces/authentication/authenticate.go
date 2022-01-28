package authentication

import (
	"context"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/types/errors"
)

func (u *workspace) Authenticate(ctx context.Context, req AuthenticationRequest) (access.Access, error) {
	const operation = "Workspace.Authentication.Authentication"
	callParams := errors.AdditionalData{Key: "request", Value: req}

	if err := req.Document.Validate(); err != nil {
		return access.Access{}, errors.Wrap(err, operation, callParams)
	}

	acc, err := u.ac.GetWithCPF(ctx, req.Document)
	if err != nil {
		return access.Access{}, errors.Wrap(err, operation, callParams)
	}

	if err := acc.Secret.Compare(req.Secret); err != nil {
		return access.Access{}, errors.Wrap(err, operation, callParams)
	}

	tok, err := u.tk.Create(acc.ExternalID)
	if err != nil {
		return access.Access{}, errors.Wrap(err, operation, callParams)
	}

	return tok, nil

}
