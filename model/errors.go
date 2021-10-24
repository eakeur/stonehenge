package model

import "errors"

var (
	ErrCreating       = errors.New("ops! an error occurred while adding your stuff. please, try again later")
	ErrAccountInvalid = errors.New("ops! the account id provided is invalid. please, try again with a valid one")
	ErrLoginInvalid   = errors.New("ops! the cpf provided does not have an account yet. come on, register yourself, you'll love it")
	ErrInvalidBody    = errors.New("ops! it seems like the json provided in the request body is invalid. please check if you typed everything correctly")
	ErrCPFInvalid     = errors.New("ops! the cpf provided is invalid. please, try again with a valid cpf")
	ErrUnauthorized   = errors.New("ops! you don't have access to this resource. please, log in with an account that has it")
	ErrForbidden      = errors.New("ops! to access this resource, you have to be logged in. please, log in or create a new account")
	ErrWrongPassword  = errors.New("ops! wrong password. please check if you typed everything right")
	ErrPostTransfer   = errors.New("ops! an error occurred while transfering your money. please be sure if the amount has left your account or contact us")
	ErrNoMoney        = errors.New("account does not have money enough to withdraw")
	ErrAmountInvalid  = errors.New("ops! the value you requested to transfer is invalid. a transfer must have a value bigger than 0")
	ErrAccountExists  = errors.New("ops! it seems like there is already an account with this cpf. please, try logging in")
	ErrSameTransfer   = errors.New("ops! it seems like you want to transfer money to yourself. you know that that's not gonna make you rich")
	ErrInternal       = errors.New("ops! an internal error occurred while processing your request. please, try again later")
	ErrNotFound       = errors.New("ops! We could not find the requested item or it does not exist. what about checking if you typed everything right?")
)
