package accounts

import (
	"context"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
)

type AuthenticationRequest struct {
	Document document.Document
	Secret password.Password
}

func (u *workspace) Authenticate(ctx context.Context, req AuthenticationRequest) (*id.ID, error) {
	if err := req.Document.Validate(); err != nil {
		return nil, err
	}

	acc, err := u.ac.GetWithCPF(ctx, req.Document)
	if err != nil {
		return nil, err
	}

	if err := acc.Secret.CompareWithString(req.Secret.Hash()); err != nil {
		return nil, err
	}

	return &acc.Id, nil

}
