package account

import (
	"context"
	"stonehenge/app/core/types/id"
)

func (u *workspace) Authenticate(ctx context.Context, req AuthenticationRequest) (id.External, error) {
	if err := req.Document.Validate(); err != nil {
		return id.Zero, err
	}

	acc, err := u.ac.GetWithCPF(ctx, req.Document)
	if err != nil {
		return id.Zero, err
	}

	if err := acc.Secret.Compare(req.Secret); err != nil {
		return id.Zero, err
	}

	return acc.ExternalID, nil

}
