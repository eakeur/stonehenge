package accounts

import (
	"net/http"
	"stonehenge/app/core/types/document"
	"stonehenge/app/gateway/api/accounts/schema"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/workspaces/account"
)

type LoginRequestBody struct {
	Document document.Document
	Secret   string
}

// Authenticate logs an applicant in
func (c *Controller) Authenticate(r *http.Request) rest.Response {
	body, err := schema.NewAuthenticationRequest(r.Body)
	if err != nil {
		return rest.BuildErrorResult(err)
	}

	ctx := r.Context()
	acc, err := c.workspace.Authenticate(ctx, account.AuthenticationRequest{
		Document: document.Document(body.Document),
		Secret:   body.Secret,
	})
	if err != nil {
		return rest.BuildErrorResult(err)
	}

	return rest.BuildOKResult(schema.AuthenticationResponse{Token: acc.Token})

}
