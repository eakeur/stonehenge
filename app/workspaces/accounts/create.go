package accounts

import (
	"context"
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"time"
)

// initialBalance expressed in cents
const initialBalance currency.Currency = 5000

type CreateInput struct {
	Document document.Document
	Secret   password.Password
	Name     string
}

type CreateOutput struct {
	AccountID id.ID
	CreatedAt time.Time
}

func (u *workspace) Create(ctx context.Context, req CreateInput) (*CreateOutput, error) {

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
		Secret:   req.Secret.Hash(),
		Document: req.Document,
		Balance:  initialBalance,
	}

	ctx, err = u.ac.StartOperation(ctx)
	if err != nil {
		//TODO create could not start operation error
		return nil, err
	}

	accountId, err := u.ac.Create(ctx, &acc)
	if err != nil {
		u.ac.RollbackOperation(ctx)
		return nil, err
	}
	err = u.ac.CommitOperation(ctx)
	if err != nil {
		//TODO create could not finish operation error
		return nil, err
	}

	return &CreateOutput{
		AccountID: *accountId,
		CreatedAt: acc.CreatedAt,
	}, nil
}
