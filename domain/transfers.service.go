package domain

import (
	model "stonehenge/model"
	"stonehenge/providers"
)

// Gets all transfer made by or to the account's owner
func GetAllTransfers(accountId string, madeToMe bool) ([]model.Transfer, error) {
	return providers.TransfersRepository.GetTransfers(accountId, madeToMe)
}

// Transfers an amount of money from an account to another
func TransferMoney(transfer model.Transfer) (*string, error) {

	// Checks if the amount is not zero or negative
	if transfer.Amount <= 0 {
		return nil, model.ErrAmountInvalid
	}

	if transfer.AccountDestinationId == transfer.AccountOriginId {
		return nil, model.ErrSameTransfer
	}

	origin, destination, err := getApplicantsAccounts(transfer)
	if err != nil {
		return nil, err
	}

	// Tries to widthdraw the origin account's money, if there is any
	_, withdrawalError := origin.Withdraw(transfer.Amount)
	if withdrawalError != nil {
		return nil, withdrawalError
	}
	// Deposits the amount to the destination account
	destination.Deposit(transfer.Amount)

	updateErr := updateRequestedAccounts(origin, destination)
	if updateErr != nil {
		return nil, updateErr
	}

	id, finalErr := providers.TransfersRepository.AddTransfer(transfer)
	if finalErr != nil {
		return nil, model.ErrPostTransfer
	}

	return &id, nil
}

// Update accounts after transfer
func updateRequestedAccounts(origin *model.Account, destination *model.Account) error {
	originUpdateError := providers.AccountsRepository.UpdateAccount(origin.Id, *origin)
	if originUpdateError != nil {
		return model.ErrPostTransfer
	}

	destinationUpdateError := providers.AccountsRepository.UpdateAccount(destination.Id, *destination)
	if destinationUpdateError != nil {
		return model.ErrPostTransfer
	}

	return nil
}

// Fetches the applicants accounts
func getApplicantsAccounts(transfer model.Transfer) (*model.Account, *model.Account, error) {
	// Gets origin account and verifies if it exists
	origin, err := GetAccountById(transfer.AccountOriginId)
	if err != nil {
		return nil, nil, model.ErrAccountInvalid
	}

	// Gets the destination account and verifies if it exists
	destination, err := GetAccountById(transfer.AccountDestinationId)
	if err != nil {
		return nil, nil, model.ErrAccountInvalid
	}

	return origin, destination, nil
}
