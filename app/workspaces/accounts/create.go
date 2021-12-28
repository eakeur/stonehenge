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
	AccountID id.ExternalID
	CreatedAt time.Time
}

func (u *workspace) Create(ctx context.Context, req CreateInput) (CreateOutput, error) {

	response := CreateOutput{}

	// Checks document's consistency
	if err := req.Document.Validate(); err != nil {
		return response, err
	}

	//Checks document uniqueness
	err := u.ac.CheckExistence(ctx, req.Document)
	if err != nil {
		return response, err
	}

	acc := account.Account{
		Name:     req.Name,
		Secret:   req.Secret.Hash(),
		Document: req.Document,
		Balance:  initialBalance,
	}

	ctx, err = u.ac.StartOperation(ctx)
	if err != nil {
		return response, account.ErrCreating
	}

	accountId, err := u.ac.Create(ctx, &acc)
	if err != nil {
		u.ac.RollbackOperation(ctx)
		return response, account.ErrCreating
	}
	err = u.ac.CommitOperation(ctx)
	if err != nil {
		return response, account.ErrCreating
	}

	response.AccountID = accountId
	response.CreatedAt = acc.CreatedAt

	return response, nil
}
