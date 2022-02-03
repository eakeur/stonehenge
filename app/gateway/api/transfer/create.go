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

// Create godoc
// @Summary      Creates transfer
// @Description  Creates a transfer with values specified on body
// @Tags         Transfers
// @Param        transfer body schema.CreateTransferRequest true "Transfer info"
// @Accept       json
// @Produce      json
// @Success      201  {object}  schema.CreateTransferResponse
// @Failure      400  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Security     AuthKey
// @Router       /transfers [post]
func (c *controller) Create(r *http.Request) rest.Response {
	const operation = "Controller.Transfer.Create"
	ctx := r.Context()
	body, err := schema.NewCreateTransferRequest(r.Body)
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
	return c.builder.BuildCreatedResult(schema.CreateTransferResponse{
		RemainingBalance: result.Output.RemainingBalance.ToStandardCurrency(),
		TransferID:       result.Output.TransferId.String(),
	}).WithSuccessLog(ctx, "transfer created successfully")
}
