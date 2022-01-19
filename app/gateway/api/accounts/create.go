package accounts

import (
	"encoding/json"
	"io"
	"net/http"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/password"
	"stonehenge/app/gateway/api/responses"
	"stonehenge/app/workspaces/account"
)

type PostRequestBody struct {
	Document string
	Secret   string
	Name     string
}

// Create creates a new account with the data passed in
func (c *Controller) Create(rw http.ResponseWriter, r *http.Request) {
	body, err := getPostRequestBody(r.Body)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusBadRequest, err)
		return
	}
	req := account.CreateInput{
		Document: document.Document(body.Document),
		Secret:   password.Password(body.Secret),
		Name:     body.Name,
	}

	acc, err := c.workspace.Create(r.Context(), req)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusBadRequest, err)
		return
	}

	rw.Header().Add("Authorization", "Bearer "+ acc.Access.Token)

	rw.WriteHeader(http.StatusCreated)

}

func getPostRequestBody(body io.ReadCloser) (PostRequestBody, error) {
	defer body.Close()
	req := PostRequestBody{}
	err := json.NewDecoder(body).Decode(&req)
	if err != nil {
		return req, err
	}

	if req.Document == "" || req.Secret == "" || req.Name == "" {
		return req, err
	}

	return req, err
}
