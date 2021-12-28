package accounts

import (
	"context"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
)

type AuthenticationRequest struct {
	Document document.Document
	Secret   password.Password
}

func (u *workspace) Authenticate(ctx context.Context, req AuthenticationRequest) (id.ExternalID, error) {
	if err := req.Document.Validate(); err != nil {
		return "", err
	}

	acc, err := u.ac.GetWithCPF(ctx, req.Document)
	if err != nil {
		return "", err
	}

	if err := acc.Secret.CompareWithString(req.Secret.Hash()); err != nil {
		return "", err
	}

	return acc.ExternalID, nil

}
