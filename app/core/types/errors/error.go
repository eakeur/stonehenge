package errors

import "fmt"

type AdditionalData struct {
	Key   string
	Value interface{}
}

type Error struct {
	operation   string
	err         error
	parameters  []AdditionalData
	parentError *Error
}

func (e Error) Error() string {
	return fmt.Sprintf("message: %s | operation: %s", e.err.Error(), e.operation)
}

func Wrap(err error, operation string, parameters ...AdditionalData) *Error {
	if err == nil {
		return nil
	}

	domainError := Error{
		parameters: parameters,
		operation:  operation,
		err:        err,
	}

	if typedError, ok := err.(*Error); ok {
		typedError.parentError = &domainError
	}

	return &domainError
}
