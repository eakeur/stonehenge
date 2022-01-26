package schema

import (
	"encoding/json"
	"io"
)

type AuthenticationRequest struct {
	// Document is the CPF of the applicant. Must be numbers only
	Document string `json:"cpf" example:"23100299900"`

	//Secret is the password of the account the applicant wants to authenticate to
	Secret string `json:"secret"`
}

func NewAuthenticationRequest(body io.ReadCloser) (AuthenticationRequest, error) {
	defer body.Close()
	req := AuthenticationRequest{}
	err := json.NewDecoder(body).Decode(&req)
	if err != nil {
		return AuthenticationRequest{}, err
	}
	return req, nil
}
