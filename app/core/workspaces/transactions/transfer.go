package transactions

import (
	"context"
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/model/transfer"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"time"
)

type TransferRequest struct {
	OriginId id.ID
	DestId   id.ID
	Amount   currency.Currency
}

func (u *workspace) Transfer(ctx context.Context, req TransferRequest) error {
	t := &transfer.Transfer{
		OriginId:      req.OriginId,
		DestinationId: req.DestId,
		Amount:        req.Amount,
	}

	// Checks if the amount is valid (bigger than 0)
	if t.Amount <= 0 {
		return transfer.ErrAmountInvalid
	}

	// Checks if the origin and destination accounts are the same
	if t.DestinationId == t.OriginId {
		return transfer.ErrSameAccount
	}

	// Fetches the origin account and checks for errors
	origin, err := u.ac.Get(ctx, req.OriginId)
	if err != nil {
		return account.ErrNotFound
	}

	// Validates if the balance of the origin is zero and if it's sufficient to accomplish the operation
	if origin.Balance <= 0 || req.Amount > origin.Balance {
		return account.ErrNoMoney
	}

	// Fetches the origin account and checks for errors
	dest, err := u.ac.Get(ctx, req.DestId)
	if err != nil {
		return account.ErrNotFound
	}

	ctx, err = u.ac.StartOperation(ctx)
	if err != nil {
		//TODO create could not start operation error
		return err
	}

	// Updates the balance of the origin account after transaction
	err = u.ac.UpdateBalance(ctx, req.OriginId, origin.Balance-req.Amount)
	if err != nil {
		return err
	}

	// Updates the balance of the destination account after transaction
	err = u.ac.UpdateBalance(ctx, req.DestId, dest.Balance+req.Amount)
	if err != nil {
		return err
	}

	t.EffectiveDate = time.Now()
	_, err = u.tr.Create(ctx, t)
	if err != nil {
		return err
	}

	err = u.ac.FinishOperation(ctx)
	if err != nil {
		//TODO create could not finish operation error
		return err
	}

	return nil
}
