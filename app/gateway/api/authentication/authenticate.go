package authentication

import (
	"net/http"
	"stonehenge/app/core/types/document"
	"stonehenge/app/gateway/api/authentication/schema"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/workspaces/authentication"
)

type LoginRequestBody struct {
	Document document.Document
	Secret   string
}

// Authenticate logs an applicant in
func (c *Controller) Authenticate(r *http.Request) rest.Response {
	const operation = "Controller.Authentication.Authenticate"
	ctx := r.Context()
	body, err := schema.NewAuthenticationRequest(r.Body)
	if err != nil {
		c.logger.Error(ctx, operation, err.Error())
		return rest.BuildErrorResult(err)
	}

	acc, err := c.workspace.Authenticate(ctx, authentication.AuthenticationRequest{
		Document: document.Document(body.Document),
		Secret:   body.Secret,
	})
	if err != nil {
		c.logger.Error(ctx, operation, err.Error())
		return rest.BuildErrorResult(err)
	}

	c.logger.Trace(ctx, operation, "finished process successfully")
	return rest.
		BuildOKResult(schema.AuthenticationResponse{Token: acc.Token}).
		AddHeaders("Authorization", "Bearer "+acc.Token)
}
