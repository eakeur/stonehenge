package transfers

import (
	"encoding/json"
	"io"
	"net/http"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/http/common"
	"stonehenge/app/workspaces/transfers"
)

type PostRequestBody struct {
	DestinationId string
	Amount        int
}

func (c *Controller) Create(rw http.ResponseWriter, r *http.Request) {
	body := getBody(r.Body)

	accountID := common.FetchContextUser(r.Context())

	req := transfers.CreateInput{
		OriginId: accountID,
		DestId:   id.ID(body.DestinationId),
		Amount:   currency.Currency(body.Amount),
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

	if req.DestinationId == "" || req.Amount == 0 {
		return nil
	}

	return &req
}
