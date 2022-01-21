package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	// HTTPStatus points out the type of result the client must expect, as informed by the HTTP patterns
	HTTPStatus int `json:"http_status,omitempty"`

	// Error informs the client about what exactly happened in case their request fails
	Error error `json:"error,omitempty"`

	// Content holds the payload of the request, which is the most valuable information for the client
	Content interface{} `json:"content,omitempty"`

	// Headers stores information to be set on header
	Headers map[string]string `json:"-"`
}

type Error struct {
	// Status is the HTTP status related to this error
	Status int `json:"-"`

	// Code is a unique identifier of this error
	Code string `json:"code,omitempty"`

	// Message is a description of this error context
	Message string `json:"message,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func WriteSuccessfulJSON(w http.ResponseWriter, status int, content interface{}) error {
	if content != nil {
		body, err := json.Marshal(content)
		if err != nil {
			return err
		}
		_, err = w.Write(body)
		if err != nil {
			return err
		}
	}
	w.WriteHeader(status)
	return nil
}

func WriteErrorResponse(w http.ResponseWriter, status int, message error) {
	w.WriteHeader(status)
	_, err := w.Write([]byte(message.Error()))
	if err != nil {
		return
	}
}
