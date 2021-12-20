package accounts

import (
	"encoding/json"
	"io"
	"net/http"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/password"
	"stonehenge/app/gateway/api/common"
	"stonehenge/app/workspaces/accounts"
)

type LoginRequestBody struct {
	Document document.Document
	Secret   password.Password
}

// Authenticate logs an applicant in
func (c *Controller) Authenticate(rw http.ResponseWriter, r *http.Request) {
	body := getAuthBody(r.Body)
	ctx := r.Context()
	req := accounts.AuthenticationRequest{
		Document: body.Document,
		Secret:   body.Secret,
	}
	id, err := c.workspace.Authenticate(ctx, req)
	if err != nil {
		return
		// TODO implement error
	}

	tok, err := common.CreateToken(*id)
	if err != nil {
		return
		// TODO err
	}

	common.AssignToken(rw, tok)

}

func getAuthBody(body io.ReadCloser) *LoginRequestBody {
	defer body.Close()
	req := LoginRequestBody{}
	err := json.NewDecoder(body).Decode(&req)
	if err != nil {
		return nil
	}

	if req.Document == "" || req.Secret == "" {
		return nil
	}

	return &req
}
