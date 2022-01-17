package account

import (
	"context"
	"stonehenge/app/core/entities/access"
)

func (u *workspace) Authenticate(ctx context.Context, req AuthenticationRequest) (access.Access, error) {
	if err := req.Document.Validate(); err != nil {
		return access.Access{}, err
	}

	acc, err := u.ac.GetWithCPF(ctx, req.Document)
	if err != nil {
		return access.Access{}, err
	}

	if err := acc.Secret.Compare(req.Secret); err != nil {
		return access.Access{}, err
	}

	tok, err := u.tk.Create(acc.ExternalID)
	if err != nil {
		return access.Access{}, err
	}

	return tok, nil

}
