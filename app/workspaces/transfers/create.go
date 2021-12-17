package transfers

import (
	"context"
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/model/transfer"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"time"
)

type CreateInput struct {
	OriginId id.ID
	DestId   id.ID
	Amount   currency.Currency
}

type CreateOutput struct {
	RemainingBalance currency.Currency
	TransferId id.ID
	CreatedAt time.Time
}

func (u *workspace) Create(ctx context.Context, req CreateInput) (*CreateOutput, error) {
	t := &transfer.Transfer{
		OriginId:      req.OriginId,
		DestinationId: req.DestId,
		Amount:        req.Amount,
	}

	// Checks if the amount is valid (bigger than 0)
	if t.Amount <= 0 {
		return nil, transfer.ErrAmountInvalid
	}

	// Checks if the origin and destination accounts are the same
	if t.DestinationId == t.OriginId {
		return nil, transfer.ErrSameAccount
	}

	// Fetches the origin account and checks for errors
	origin, err := u.ac.Get(ctx, req.OriginId)
	if err != nil {
		return nil, account.ErrNotFound
	}

	// Validates if the balance of the origin is zero and if it's sufficient to accomplish the operation
	if origin.Balance <= 0 || req.Amount > origin.Balance {
		return nil, account.ErrNoMoney
	}

	// Fetches the origin account and checks for errors
	dest, err := u.ac.Get(ctx, req.DestId)
	if err != nil {
		return nil, account.ErrNotFound
	}

	ctx, err = u.ac.StartOperation(ctx)
	if err != nil {
		//TODO create could not start operation error
		return nil, err
	}

	// Updates the balance of the origin account after transaction
	remaining := origin.Balance-req.Amount
	err = u.ac.UpdateBalance(ctx, req.OriginId, remaining)
	if err != nil {
		u.ac.RollbackOperation(ctx)
		return nil, err
	}

	// Updates the balance of the destination account after transaction
	err = u.ac.UpdateBalance(ctx, req.DestId, dest.Balance+req.Amount)
	if err != nil {
		u.ac.RollbackOperation(ctx)
		return nil, err
	}

	// Creates a transfer register on storage
	t.EffectiveDate = time.Now()
	transferId, err := u.tr.Create(ctx, t)
	if err != nil {
		u.ac.RollbackOperation(ctx)
		return nil, err
	}

	err = u.ac.CommitOperation(ctx)
	if err != nil {
		u.ac.RollbackOperation(ctx)
		//TODO create could not finish operation error
		return nil, err
	}

	return &CreateOutput{
		RemainingBalance: remaining,
		TransferId: *transferId,
		CreatedAt: t.CreatedAt,
	}, nil
}
