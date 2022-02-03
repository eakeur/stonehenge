package schema

import (
	"encoding/json"
	"io"
)

type CreateAccountRequest struct {
	// Document is the applicant's CPF. Must be numbers only
	Document string `json:"cpf" example:"23100299900"`

	// Secret is the account's password. It will be used to authenticate afterwards
	Secret string `json:"secret"`

	// Name is the applicant's full name
	Name string `json:"name" example:"Alan Turing"`
}

func NewCreateAccountRequest(body io.ReadCloser) (CreateAccountRequest, error) {
	defer body.Close()
	var req CreateAccountRequest
	err := json.NewDecoder(body).Decode(&req)
	if err != nil {
		return CreateAccountRequest{}, err
	}
	return req, nil
}
