package transfer

import (
	"fmt"
	"net/http"
	"net/url"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/erring"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/transfer/schema"
	"time"
)

// List godoc
// @Summary      List transfers
// @Description  List all transfers that match the given filter
// @Tags         Transfers
// @Param       originID  query string  false  "Account origin id"
// @Param       destinationID  query string  false  "Account destination id"
// @Param       made_since  query string  false  "Initial date"
// @Param       made_until  query string  false  "Final date"
// @Produce      json
// @Success      200  {object}  []schema.ListTransferResponse
// @Failure      400  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Security     AuthKey
// @Router       /api/v1/transfers [get]
func (c *controller) List(r *http.Request) rest.Response {
	const operation = "Controller.Transfer.List"
	ctx := r.Context()
	filters := filter(r.URL.Query())
	list, err := c.workspace.List(ctx, filters)
	if err != nil {
		err = erring.Wrap(err, operation)
		return c.builder.BuildErrorResult(err).WithErrorLog(ctx)
	}
	length := len(list)
	res := make([]schema.ListTransferResponse, length)
	for i, ref := range list {
		res[i] = schema.ListTransferResponse{
			ExternalID:    ref.ExternalID.String(),
			OriginID:      ref.Details.OriginExternalID.String(),
			DestinationID: ref.Details.DestinationExternalID.String(),
			Amount:        ref.Amount.ToStandardCurrency(),
			EffectiveDate: ref.EffectiveDate,
		}
	}
	return c.builder.BuildOKResult(res).
		AddHeaders("X-Total-Count", fmt.Sprint(length)).
		WithSuccessLog(ctx, fmt.Sprintf("listed accounts with %d results", length))
}

func filter(values url.Values) transfer.Filter {
	f := transfer.Filter{}

	if ori := id.ExternalFrom(values.Get("originID")); ori != id.Zero {
		f.OriginID = ori
	}

	if dest := id.ExternalFrom(values.Get("destinationID")); dest != id.Zero {
		f.DestinationID = dest
	}

	if ini := values.Get("made_since"); ini != "" {
		date, err := time.Parse("2006-01-02", ini)
		if err == nil {
			f.InitialDate = date
		}
	}

	if fin := values.Get("made_until"); fin != "" {
		date, err := time.Parse("2006-01-02", fin)
		if err == nil {
			f.FinalDate = date
		}
	}

	return f
}
