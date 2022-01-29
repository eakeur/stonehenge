package account

import (
	"fmt"
	"net/http"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/erring"
	"stonehenge/app/core/types/password"
	"stonehenge/app/gateway/api/account/schema"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/workspaces/account"
)

// Create creates a new account with the data passed in
func (c *controller) Create(r *http.Request) rest.Response {
	const operation = "Controller.Account.Create"
	ctx := r.Context()

	req, err := schema.NewCreateRequest(r.Body)
	if err != nil {
		err = erring.Wrap(err, operation)
		return c.builder.BuildErrorResult(err).WithErrorLog(ctx)
	}

	input := account.CreateInput{
		Document: document.Document(req.Document),
		Secret:   password.From(req.Secret),
		Name:     req.Name,
	}

	acc, err := c.workspace.Create(ctx, input)
	if err != nil {
		err = erring.Wrap(err, operation)
		return c.builder.BuildErrorResult(err).WithErrorLog(ctx)
	}

	return c.builder.
		BuildCreatedResult(schema.CreateResponse{AccountID: acc.AccountID.String(), Token: acc.Access.Token}).
		AddHeaders("Authorization", "Bearer "+acc.Access.Token).
		WithSuccessLog(ctx, fmt.Sprintf("account created successfully with id %s", acc.Access.AccountID.String()))

}
