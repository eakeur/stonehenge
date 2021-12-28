package accounts

import (
	"net/http"
	"net/url"
	"stonehenge/app/core/model/account"
	"stonehenge/app/gateway/api/responses"
)

// List gets all accounts that satisfy the filter passed
func (c *Controller) List(rw http.ResponseWriter, r *http.Request) {
	filters := getFilter(r.URL.Query())
	list, err := c.workspace.List(r.Context(), filters)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusBadRequest, err)
		return
	}

	err = responses.WriteSuccessfulJSON(rw, http.StatusOK, list)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusInternalServerError, err)
		return
	}

}

func getFilter(values url.Values) account.Filter {
	filter := account.Filter{}

	if name := values.Get("name"); name != "" {
		filter.Name = name
	}

	return filter
}
