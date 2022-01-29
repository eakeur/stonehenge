package rest

import (
	"errors"
	"net/http"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
)

// Account errors
var (
	ErrAccountNotFound = Error{
		Status:  http.StatusNotFound,
		Code:    "stonehenge:account_not_found",
		Message: "The account with the ID informed could not be found. Please verify if you informed the right value.",
	}

	ErrAccountAlreadyExists = Error{
		Status:  http.StatusBadRequest,
		Code:    "stonehenge:account_already_exists",
		Message: "Ops! The CPF informed is already assigned to another account. Please try logging in.",
	}

	ErrAccountCannotAccess = Error{
		Status:  http.StatusForbidden,
		Code:    "stonehenge:account_cannot_access",
		Message: "Ops! The current logged in account is not allowed to access this resource. Please log in with an account that is.",
	}

	ErrAccountNoMoney = Error{
		Status:  http.StatusBadRequest,
		Code:    "stonehenge:account_no_money",
		Message: "The account with the ID informed does have enough money to complete this operation",
	}
)

// Transfer errors
var (
	ErrTransferSameAccount = Error{
		Status:  http.StatusBadRequest,
		Code:    "stonehenge:transfer_same_account",
		Message: "The origin and the destination account are the same. Please verify if you've informed the correct IDs.",
	}

	ErrTransferAmountInvalid = Error{
		Status:  http.StatusBadRequest,
		Code:    "stonehenge:account_amount_invalid",
		Message: "The amount informed to be transfer is smaller or equal to zero.",
	}

	ErrTransferNonexistentOrigin = Error{
		Status:  http.StatusBadRequest,
		Code:    "stonehenge:transfer_nonexistent_origin",
		Message: "The origin account informed does not exist or could not be found. Please choose another one.",
	}

	ErrTransferNonExistentDestination = Error{
		Status:  http.StatusBadRequest,
		Code:    "stonehenge:transfer_nonexistent_destination",
		Message: "The destination account informed does not exist or could not be found. Please choose another one.",
	}
)

// Access errors
var (
	ErrAccessNonexistent = Error{
		Status:  http.StatusUnauthorized,
		Code:    "stonehenge:access_nonexistent",
		Message: "This request should be made with a logged in user. Please authenticate",
	}

	ErrAccessExpired = Error{
		Status:  http.StatusUnauthorized,
		Code:    "stonehenge:access_unauthorized",
		Message: "Your session just expired. Please try logging in again",
	}
)

// Types errors
var (
	ErrDocumentInvalid = Error{
		Status:  http.StatusBadRequest,
		Code:    "stonehenge:document_invalid",
		Message: "The CPF provided in this request is not valid. Please verify or provide another one.",
	}

	ErrIDInvalid = Error{
		Status:  http.StatusBadRequest,
		Code:    "stonehenge:id_invalid",
		Message: "The UUID provided in this request is not a valid one",
	}

	ErrPasswordInvalid = Error{
		Status:  http.StatusUnauthorized,
		Code:    "stonehenge:password_invalid",
		Message: "The password provided is invalid. Please provide another one",
	}
)

var ErrGeneral = Error{
	Status:  http.StatusInternalServerError,
	Code:    "stonehenge:unhandled_error",
	Message: "Ops! An error we were not expecting occurred while processing your request. Please try again in a few minutes or contact us.",
}

var mappedErrors = map[error]Error{
	account.ErrNotFound:                ErrAccountNotFound,
	account.ErrAlreadyExist:            ErrAccountAlreadyExists,
	account.ErrNoMoney:                 ErrAccountNoMoney,
	account.ErrCannotAccess:            ErrAccountCannotAccess,
	transfer.ErrSameAccount:            ErrTransferSameAccount,
	transfer.ErrAmountInvalid:          ErrTransferAmountInvalid,
	transfer.ErrNonexistentOrigin:      ErrTransferNonexistentOrigin,
	transfer.ErrNonexistentDestination: ErrTransferNonExistentDestination,
	access.ErrNoAccessInContext:        ErrAccessNonexistent,
	access.ErrTokenInvalidOrExpired:    ErrAccessExpired,
	document.ErrInvalidDocument:        ErrDocumentInvalid,
	id.ErrInvalidID:                    ErrIDInvalid,
	password.ErrWrongPassword:          ErrPasswordInvalid,
}

func FindMatchingDomainError(err error) Error {
	for errKey, errValue := range mappedErrors {
		if errors.Is(err, errKey) {
			return errValue
		}
	}
	return ErrGeneral
}
