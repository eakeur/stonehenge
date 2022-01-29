package erring

import (
	"encoding/json"
	"fmt"
)

type AdditionalData struct {
	Key   string
	Value interface{}
}

type Error struct {
	Operation  string
	Err        error
	Parameters []AdditionalData
	ParentErr  *Error
}

func (e Error) Error() string {

	var parameters string
	if e.Parameters != nil {
		prms := map[string]interface{}{}
		for _, prm := range e.Parameters {
			prms[prm.Key] = prm.Value
		}

		res, err := json.Marshal(prms)
		if err == nil {
			parameters = string(res)
			parameters = parameters[1 : len(parameters)-1]
		}

		return fmt.Sprintf("%s with %s -> %s", e.Operation, parameters, e.Err.Error())
	}

	return fmt.Sprintf("%s -> %s", e.Operation, e.Err.Error())
}

func (e Error) Unwrap() error {
	return e.Err
}

func Wrap(err error, operation string, parameters ...AdditionalData) *Error {
	if err == nil {
		return nil
	}

	domainError := Error{
		Parameters: parameters,
		Operation:  operation,
		Err:        err,
	}

	if typedError, ok := err.(*Error); ok {
		typedError.ParentErr = &domainError
	}

	return &domainError
}
