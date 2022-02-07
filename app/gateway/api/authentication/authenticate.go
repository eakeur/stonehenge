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

// Authenticate godoc
// @Summary      Authenticates account
// @Description  Authenticates an account with its credentials
// @Tags         Login
// @Param        account body schema.AuthenticationRequest true "Account info"
// @Accept       json
// @Produce      json
// @Success      201  {object}  schema.AuthenticationResponse
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /api/v1/login [post]
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
