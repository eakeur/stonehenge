package schema

import "io"

type CreateRequest struct {
	// Document is the applicant's CPF. Must be numbers only
	Document string `json:"cpf" example:"23100299900"`

	// Secret is the account's password. It will be used to authenticate afterwards
	Secret string `json:"secret"`

	// Name is the applicant's full name
	Name string `json:"name" example:"Alan Turing"`
}

func NewCreateRequest(body io.ReadCloser) (CreateRequest, error){
	return CreateRequest{}, nil
}