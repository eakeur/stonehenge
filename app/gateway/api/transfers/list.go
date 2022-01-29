package transfers

import (
	"net/http"
	"net/url"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/transfers/schema"
	"time"
)

// List gets all transfers of this actual account
func (c controller) List(r *http.Request) rest.Response {
	const operation = "Controller.Transfer.Create"
	ctx := r.Context()
	filters := filter(r.URL.Query())
	list, err := c.workspace.List(ctx, filters)
	if err != nil {
		c.logger.Error(ctx, operation, err.Error())
		return rest.BuildErrorResult(err)
	}
	res := make([]schema.ListResponse, len(list))
	for i, ref := range list {
		res[i] = schema.ListResponse{
			ExternalID: ref.ExternalID.String(),
			// TODO change the IDs below to type External
			//OriginID:      ref.OriginID.String(),
			//DestinationID: ref.DestinationID.String(),
			Amount:        ref.Amount.ToStandardCurrency(),
			EffectiveDate: ref.EffectiveDate,
		}
	}
	c.logger.Trace(ctx, operation, "finished process successfully")
	return rest.BuildOKResult(list)
}

func filter(values url.Values) transfer.Filter {
	f := transfer.Filter{}

	if ori := id.ExternalFrom(values.Get("origin")); ori != id.Zero {
		f.OriginID = ori
	}

	if dest := id.ExternalFrom(values.Get("destination")); dest != id.Zero {
		f.OriginID = dest
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
			f.InitialDate = date
		}
	}

	return f
}
