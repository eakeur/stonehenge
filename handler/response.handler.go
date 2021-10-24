package handler

import (
	"net/http"
	model "stonehenge/model"
)

// Writes the response body to the object passed in
func SendResponse(response http.ResponseWriter, body []byte, status int) {
	response.WriteHeader(status)
	response.Write(body)
}

// Writes an error message to the response and sets the status code based on the very same error
func SendErrorResponse(response http.ResponseWriter, body error) {
	response.WriteHeader(GetStatusCodeByError(body))
	response.Write([]byte(body.Error()))
}

// Gets the status code based on the error model passed as parameter
func GetStatusCodeByError(err error) int {
	switch err {
	case model.ErrCreating:
		return http.StatusInternalServerError
	case model.ErrAccountInvalid:
		return http.StatusBadRequest
	case model.ErrLoginInvalid:
		return http.StatusUnauthorized
	case model.ErrInvalidBody:
		return http.StatusBadRequest
	case model.ErrCPFInvalid:
		return http.StatusBadRequest
	case model.ErrUnauthorized:
		return http.StatusUnauthorized
	case model.ErrForbidden:
		return http.StatusForbidden
	case model.ErrWrongPassword:
		return http.StatusUnauthorized
	case model.ErrPostTransfer:
		return http.StatusInternalServerError
	case model.ErrNoMoney:
		return http.StatusBadRequest
	case model.ErrAmountInvalid:
		return http.StatusBadRequest
	case model.ErrAccountExists:
		return http.StatusBadRequest
	case model.ErrInternal:
		return http.StatusBadRequest
	case model.ErrSameTransfer:
		return http.StatusBadRequest
	case model.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
