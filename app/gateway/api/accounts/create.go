package accounts

import (
	"net/http"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/password"
	"stonehenge/app/gateway/api/accounts/schema"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/workspaces/account"
)

// Create creates a new account with the data passed in
func (c *Controller) Create(r *http.Request) rest.Response {
	const operation = "Controller.Account.Create"
	ctx := r.Context()
	req, err := schema.NewCreateRequest(r.Body)
	if err != nil {
		c.logger.Error(ctx, operation, err.Error())
		return rest.BuildBadRequestResult(err)
	}

	input := account.CreateInput{
		Document: document.Document(req.Document),
		Secret:   password.Password(req.Secret),
		Name:     req.Name,
	}

	acc, err := c.workspace.Create(r.Context(), input)
	if err != nil {
		c.logger.Error(ctx, operation, err.Error())
		return rest.BuildErrorResult(err)
	}

	c.logger.Trace(ctx, operation, "finished process successfully")
	return rest.
		BuildCreatedResult(schema.CreateResponse{AccountID: acc.AccountID.String(), Token: acc.Access.Token}).
		AddHeaders("Authorization", "Bearer "+acc.Access.Token)

}
