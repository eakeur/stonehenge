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
	const operation = "Controller.Transfer.Create"
	ctx := r.Context()
	body, err := schema.NewCreateRequest(r.Body)
	if err != nil {
		c.logger.Error(ctx, operation, err.Error())
		return rest.BuildErrorResult(err)
	}

	create, err := c.workspace.Create(ctx, transfer.CreateInput{
		DestID: id.ExternalFrom(body.DestinationID),
		Amount: currency.FromStandardCurrency(body.Amount),
	})
	if err != nil {
		c.logger.Error(ctx, operation, err.Error())
		return rest.BuildErrorResult(err)
	}
	c.logger.Trace(ctx, operation, "finished process successfully")
	return rest.BuildCreatedResult(schema.CreateResponse{
		RemainingBalance: create.RemainingBalance.ToStandardCurrency(),
		TransferID:       create.TransferId.String(),
	})
}
