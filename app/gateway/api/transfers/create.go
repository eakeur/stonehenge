package transfers

import (
	"net/http"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/transfers/schema"
	"stonehenge/app/workspaces/transfer"
)

func (c controller) Create(r *http.Request) rest.Response {
	body, err := schema.NewCreateRequest(r.Body)
	if err != nil {
		return rest.BuildErrorResult(err)
	}

	create, err := c.workspace.Create(r.Context(), transfer.CreateInput{
		DestID: id.ExternalFrom(body.DestinationID),
		Amount: currency.FromStandardCurrency(body.Amount),
	})
	if err != nil {
		return rest.BuildErrorResult(err)
	}

	return rest.BuildCreatedResult(schema.CreateResponse{
		RemainingBalance: create.RemainingBalance.ToStandardCurrency(),
		TransferID:       create.TransferId.String(),
	})
}
