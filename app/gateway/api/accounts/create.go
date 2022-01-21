package accounts

import (
	"net/http"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/password"
	"stonehenge/app/gateway/api/accounts/schema"
	"stonehenge/app/gateway/api/responses"
	"stonehenge/app/workspaces/account"
)

type PostRequestBody struct {
	Document string
	Secret   string
	Name     string
}

// Create creates a new account with the data passed in
func (c *Controller) Create(r *http.Request) responses.Response {
	req, err := schema.NewCreateRequest(r.Body)
	if err != nil {
		return responses.BuildBadRequestResult(err)
	}

	input := account.CreateInput{
		Document: document.Document(req.Document),
		Secret:   password.Password(req.Secret),
		Name:     req.Name,
	}

	acc, err := c.workspace.Create(r.Context(), input)
	if err != nil {
		responses.BuildErrorResult(err)
	}

	res := responses.BuildCreatedResult(schema.CreateResponse{AccountID: acc.AccountID.String(), Token: acc.Access.Token})
	return res

}
