package account

import "errors"

var (

	// ErrCreating occurs when an untracked error happens when creating a new account
	ErrCreating = errors.New("an error occurred while creating this account")

	// ErrInvalid throws when the identification provided is an invalid form of ID
	ErrInvalid = errors.New("the account id provided is invalid. please, try again with a valid one")

	// ErrNotFound expresses that the account with the identification provided does not exist, or, for some reason, could not be found
	ErrNotFound = errors.New("the account requested could not be found or does not exist")

	// ErrAlreadyExist points out that an account with the CPF provided in the request, which is unique, already exists in the database
	ErrAlreadyExist = errors.New("an account with the document provided already exists")

	// ErrNoMoney expresses that the account requesting a financial operation does not have budget enough to do so
	ErrNoMoney = errors.New("account does not have money enough for this operation")

	// ErrAmountInvalid throws when a financial operation request is made with an invalid amount of money, which is equal or less than zero
	ErrAmountInvalid = errors.New("the amount of money provided is not valid. please provide a value bigger than zero")
)
