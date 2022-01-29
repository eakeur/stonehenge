package transfer

import (
	"net/http"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/erring"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/transfer/schema"
	"stonehenge/app/workspaces/transfer"
)

func (c *controller) Create(r *http.Request) rest.Response {
	const operation = "Controller.Transfer.Create"
	ctx := r.Context()
	body, err := schema.NewCreateRequest(r.Body)
	if err != nil {
		err = erring.Wrap(err, operation)
		return c.builder.BuildErrorResult(err).WithErrorLog(ctx)
	}

	create, err := c.workspace.Create(ctx, transfer.CreateInput{
		DestID: id.ExternalFrom(body.DestinationID),
		Amount: currency.FromStandardCurrency(body.Amount),
	})
	if err != nil {
		err = erring.Wrap(err, operation)
		return c.builder.BuildErrorResult(err).WithErrorLog(ctx)
	}
	return c.builder.BuildCreatedResult(schema.CreateResponse{
		RemainingBalance: create.RemainingBalance.ToStandardCurrency(),
		TransferID:       create.TransferId.String(),
	}).WithSuccessLog(ctx, "transfer created successfully")
}
