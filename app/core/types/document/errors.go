package document

import "errors"

var (

	// ErrInvalidDocument occurs when a CPF provided by the client is invalid or contains special characters
	ErrInvalidDocument = errors.New("the document provided is invalid. please inform a valid cpf")
)
