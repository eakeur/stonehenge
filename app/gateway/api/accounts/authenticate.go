package accounts

import (
	"encoding/json"
	"io"
	"net/http"
	"stonehenge/app/core/types/document"
	"stonehenge/app/gateway/api/responses"
	"stonehenge/app/workspaces/account"
)

type LoginRequestBody struct {
	Document document.Document
	Secret   string
}

// Authenticate logs an applicant in
func (c *Controller) Authenticate(rw http.ResponseWriter, r *http.Request) {
	body, err := getLoginRequestBody(r.Body)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusBadRequest, err)
	}
	ctx := r.Context()
	req := account.AuthenticationRequest{
		Document: body.Document,
		Secret:   body.Secret,
	}
	acc, err := c.workspace.Authenticate(ctx, req)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusUnauthorized, err)
		return
	}

	rw.Header().Add("Authorization", "Bearer "+acc.Token)
	rw.WriteHeader(http.StatusOK)

}

func getLoginRequestBody(body io.ReadCloser) (LoginRequestBody, error) {
	defer body.Close()
	req := LoginRequestBody{}
	err := json.NewDecoder(body).Decode(&req)
	if err != nil {
		return req, err
	}

	if req.Document == "" || req.Secret == "" {
		return req, err
	}

	return req, nil
}
