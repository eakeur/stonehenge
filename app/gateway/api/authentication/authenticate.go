package authentication

import (
	"fmt"
	"net/http"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/erring"
	"stonehenge/app/gateway/api/authentication/schema"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/workspaces/authentication"
)

// Authenticate logs an applicant in
func (c *controller) Authenticate(r *http.Request) rest.Response {
	const operation = "Controller.Authentication.Authenticate"
	ctx := r.Context()
	body, err := schema.NewAuthenticationRequest(r.Body)
	if err != nil {
		err = erring.Wrap(err, operation)
		return c.builder.BuildErrorResult(err).WithErrorLog(ctx)
	}

	acc, err := c.workspace.Authenticate(ctx, authentication.AuthenticationRequest{
		Document: document.Document(body.Document),
		Secret:   body.Secret,
	})
	if err != nil {
		err = erring.Wrap(err, operation)
		return c.builder.BuildErrorResult(err).WithErrorLog(ctx)
	}
	return c.builder.
		BuildOKResult(schema.AuthenticationResponse{Token: acc.Token}).
		AddHeaders("Authorization", "Bearer "+acc.Token).
		WithSuccessLog(ctx, fmt.Sprintf("authenticated to API user %s", acc.AccountID.String()))
}
