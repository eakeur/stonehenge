package account

import (
	"context"
	"stonehenge/pkg/stonehenge/core/model/account"
	"stonehenge/pkg/stonehenge/core/types/document"
	"stonehenge/pkg/stonehenge/core/types/id"
	"stonehenge/pkg/stonehenge/core/types/password"
	"time"
)

type CreateInput struct {
	Document document.Document
	Secret   password.Password
	Name     string
}

type CreateOutput struct {
	AccountID id.ID
	CreatedAt time.Time
}

func (u *useCase) Create(ctx context.Context, req CreateInput) (*CreateOutput, error) {

	// Checks document's consistency
	if err := req.Document.Validate(); err != nil {
		return nil, err
	}

	//Checks document uniqueness
	err := u.ac.CheckExistence(ctx, req.Document)
	if err != nil {
		return nil, err
	}

	acc := account.Account{
		Name:     req.Name,
		Secret:   req.Secret,
		Document: req.Document,
	}

	accountId, err := u.ac.Create(ctx, &acc)
	if err != nil {
		return nil, err
	}
	return &CreateOutput{
		AccountID: *accountId,
		CreatedAt: acc.CreatedAt,
	}, nil
}
