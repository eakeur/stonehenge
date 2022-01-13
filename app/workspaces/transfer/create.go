package transfer

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"time"
)

func (u *workspace) Create(ctx context.Context, req CreateInput) (CreateOutput, error) {
	t := transfer.Transfer{
		Amount: req.Amount,
	}

	response := CreateOutput{}

	// Checks if the amount is valid (bigger than 0)
	if t.Amount <= 0 {
		return response, transfer.ErrAmountInvalid
	}

	// Checks if the origin and destination accounts are the same
	if req.DestID == req.OriginID {
		return response, transfer.ErrSameAccount
	}

	// Fetches the origin account and checks for errors
	origin, err := u.ac.GetByExternalID(ctx, req.OriginID)
	if err != nil {
		return response, transfer.ErrNonexistentOrigin
	}

	// Validates if the balance of the origin is zero and if it's sufficient to accomplish the operation
	if origin.Balance <= 0 || req.Amount > origin.Balance {
		return response, account.ErrNoMoney
	}

	// Fetches the origin account and checks for errors
	dest, err := u.ac.GetByExternalID(ctx, req.DestID)
	if err != nil {
		return response, transfer.ErrNonexistentDestination
	}

	ctx, err = u.tx.Begin(ctx)
	if err != nil {
		return response, transfer.ErrRegistering
	}
	defer u.tx.Rollback(ctx)

	t.OriginID = origin.ID
	t.DestinationID = dest.ID

	// Updates the balance of the origin account after transaction
	remaining := origin.Balance - req.Amount
	err = u.ac.UpdateBalance(ctx, req.OriginID, remaining)
	if err != nil {
		return response, err
	}

	// Updates the balance of the destination account after transaction
	err = u.ac.UpdateBalance(ctx, req.DestID, dest.Balance+req.Amount)
	if err != nil {
		return response, err
	}

	// Creates a transfer register on storage
	t.EffectiveDate = time.Now()
	t, err = u.tr.Create(ctx, t)
	if err != nil {
		return response, err
	}

	err = u.tx.Commit(ctx)
	if err != nil {
		return response, transfer.ErrRegistering
	}

	response.RemainingBalance = remaining
	response.TransferId = t.ExternalID
	response.CreatedAt = t.CreatedAt
	return response, nil
}
