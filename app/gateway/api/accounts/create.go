package accounts

import (
	"encoding/json"
	"io"
	"net/http"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/password"
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
	c.workspace.Create(r.Context(), req)

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
