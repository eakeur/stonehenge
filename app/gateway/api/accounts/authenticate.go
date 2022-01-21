package accounts

import (
	"net/http"
	"stonehenge/app/core/types/document"
	"stonehenge/app/gateway/api/accounts/schema"
	"stonehenge/app/gateway/api/responses"
	"stonehenge/app/workspaces/account"
)

type LoginRequestBody struct {
	Document document.Document
	Secret   string
}

// Authenticate logs an applicant in
func (c *Controller) Authenticate(r *http.Request) responses.Response {
	body, err := schema.NewAuthenticationRequest(r.Body)
	if err != nil {
		return responses.BuildErrorResult(err)
	}

	ctx := r.Context()
	acc, err := c.workspace.Authenticate(ctx, account.AuthenticationRequest{
		Document: document.Document(body.Document),
		Secret:   body.Secret,
	})
	if err != nil {
		return responses.BuildErrorResult(err)
	}

	return responses.BuildOKResult(schema.AuthenticationResponse{Token: acc.Token})

}