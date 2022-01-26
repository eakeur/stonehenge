package authentication

import (
	"context"
	"stonehenge/app/core/entities/access"
)

func (u *workspace) Authenticate(ctx context.Context, req AuthenticationRequest) (access.Access, error) {
	const operation = "Workspace.Authentication.Authentication"
	if err := req.Document.Validate(); err != nil {
		u.logger.Error(ctx, operation, err.Error())
		return access.Access{}, err
	}

	acc, err := u.ac.GetWithCPF(ctx, req.Document)
	if err != nil {
		u.logger.Error(ctx, operation, err.Error())
		return access.Access{}, err
	}

	if err := acc.Secret.Compare(req.Secret); err != nil {
		u.logger.Error(ctx, operation, err.Error())
		return access.Access{}, err
	}

	tok, err := u.tk.Create(acc.ExternalID)
	if err != nil {
		u.logger.Error(ctx, operation, err.Error())
		return access.Access{}, err
	}

	u.logger.Trace(ctx, operation, "finished process successfully")
	return tok, nil

}
