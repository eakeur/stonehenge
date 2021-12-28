package transfers

import (
	"encoding/json"
	"io"
	"net/http"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/common"
	"stonehenge/app/gateway/api/responses"
	"stonehenge/app/workspaces/transfers"
)

type PostRequestBody struct {
	DestinationId id.ExternalID
	Amount        int
}

func (c *Controller) Create(rw http.ResponseWriter, r *http.Request) {
	body, err := getPostRequestBody(r.Body)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusBadRequest, err)
	}

	accountID := common.FetchContextUser(r.Context())

	req := transfers.CreateInput{
		OriginId: accountID,
		DestId:   body.DestinationId,
		Amount:   currency.Currency(body.Amount),
	}
	create, err := c.workspace.Create(r.Context(), req)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusBadRequest, err)
		return
	}

	err = responses.WriteSuccessfulJSON(rw, http.StatusCreated, create)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusInternalServerError, err)
		return
	}
}

func getPostRequestBody(body io.ReadCloser) (PostRequestBody, error) {
	defer body.Close()
	req := PostRequestBody{}
	err := json.NewDecoder(body).Decode(&req)
	if err != nil {
		return req, err
	}

	if req.DestinationId == "" || req.Amount == 0 {
		return req, err
	}

	return req, nil
}
