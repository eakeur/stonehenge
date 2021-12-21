package transfer

import "errors"

var (

	// ErrRegistering occurs when a transfer fails being saved in the transfer history due to an untracked error
	ErrRegistering = errors.New("an error occurred while saving this transfer, please check your balance")

	// ErrSameAccount points out that the origin and the destination account of the transfer are equal, which must not happen
	ErrSameAccount = errors.New("the account id for the origin and destination are the same. please choose another destination")

	// ErrAmountInvalid throws when a transfer request is made with an invalid amount of money, which is equal or less than zero
	ErrAmountInvalid = errors.New("the amount of money provided is not valid. please provide a value bigger than zero")

	//ErrNotFound happens when no transfer is found
	ErrNotFound = errors.New("could not find any transfers with the given filters")
)
