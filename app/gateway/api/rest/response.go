package rest

import (
	"fmt"
)

type Response struct {
	// HTTPStatus points out the type of result the client must expect, as informed by the HTTP patterns
	HTTPStatus int `json:"http_status,omitempty"`

	// Error informs the client about what exactly happened in case their request fails
	Error error `json:"error,omitempty"`

	// Content holds the payload of the request, which is the most valuable information for the client
	Content interface{} `json:"content,omitempty"`

	// headers stores information to be set on header
	headers map[string]string
}

func (r Response) AddHeaders(key, value string) Response {
	if r.headers == nil {
		r.headers = map[string]string{}
	}
	r.headers[key] = value
	return r
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
