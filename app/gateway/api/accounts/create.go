package accounts

import (
	"encoding/json"
	"io"
	"net/http"
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/password"
	"stonehenge/app/gateway/api/common"
	"stonehenge/app/gateway/api/responses"
	"stonehenge/app/workspaces/accounts"
)

type PostRequestBody struct {
	Document string
	Secret   string
	Name     string
}

// Create creates a new account with the data passed in
func (c *Controller) Create(rw http.ResponseWriter, r *http.Request) {
	body := getBody(r.Body)
	req := accounts.CreateInput{
		Document: document.Document(body.Document),
		Secret:   password.Password(body.Secret),
		Name:     body.Name,
	}

	create, err := c.workspace.Create(r.Context(), req)
	if err != nil {
		if err == account.ErrCreating || err == account.ErrExists {
			responses.WriteErrorResponse(rw, http.StatusBadRequest, err)
		}

		responses.WriteErrorResponse(rw, http.StatusInternalServerError, err)
		return
	}

	tok, err := common.CreateToken(create.AccountID)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusInternalServerError, ErrTokenGeneration)
		return
	}

	common.AssignToken(rw, tok)

	rw.WriteHeader(http.StatusCreated)

}

func getBody(body io.ReadCloser) *PostRequestBody {
	defer body.Close()
	req := PostRequestBody{}
	err := json.NewDecoder(body).Decode(&req)
	if err != nil {
		return nil
	}

	if req.Document == "" || req.Secret == "" || req.Name == "" {
		return nil
	}

	return &req
}
