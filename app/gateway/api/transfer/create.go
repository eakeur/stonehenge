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

	waiter := c.worker.AddToQueue(ctx, transfer.CreateInput{
		DestID: id.ExternalFrom(body.DestinationID),
		Amount: currency.FromStandardCurrency(body.Amount),
	})
	result := <-waiter
	if result.Error != nil {
		err = erring.Wrap(result.Error, operation)
		return c.builder.BuildErrorResult(err).WithErrorLog(ctx)
	}
	return c.builder.BuildCreatedResult(schema.CreateResponse{
		RemainingBalance: result.Output.RemainingBalance.ToStandardCurrency(),
		TransferID:       result.Output.TransferId.String(),
	}).WithSuccessLog(ctx, "transfer created successfully")
}
