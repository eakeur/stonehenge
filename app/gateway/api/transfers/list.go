package transfers

import (
	"net/http"
	"net/url"
	"stonehenge/app/core/model/transfer"
	"stonehenge/app/gateway/api/common"
	"stonehenge/app/gateway/api/responses"
	"time"
)

// List gets all transfers of this actual account
func (c *Controller) List(rw http.ResponseWriter, r *http.Request) {
	filters := getFilter(r.URL.Query())
	filters.OriginId = string(common.FetchContextUser(r.Context()))
	list, err := c.workspace.List(r.Context(), filters)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusBadRequest, err)
	}

	err = responses.WriteSuccessfulJSON(rw, http.StatusOK, list)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusInternalServerError, err)
		return
	}
}

func getFilter(values url.Values) transfer.Filter {
	filter := transfer.Filter{}

	if ori := values.Get("origin"); ori != "" {
		filter.OriginId = ori
	}

	if dest := values.Get("destination"); dest != "" {
		filter.OriginId = dest
	}

	if ini := values.Get("made_since"); ini != "" {
		date, err := time.Parse("2006-01-02", ini)
		if err == nil {
			filter.InitialDate = date
		}
	}

	if fin := values.Get("made_until"); fin != "" {
		date, err := time.Parse("2006-01-02", fin)
		if err == nil {
			filter.InitialDate = date
		}
	}

	return filter
}
