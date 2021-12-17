package accounts

import (
	"net/http"
	"net/url"
	"stonehenge/app/core/model/account"
)

// List gets all accounts that satisfy the filter passed
func (c *Controller) List(rw http.ResponseWriter, r *http.Request) {
	filters := getFilter(r.URL.Query())
	c.workspace.List(r.Context(), filters)
}

func getFilter(values url.Values) account.Filter {
	filter := account.Filter{}

	if name := values.Get("name"); name != "" {
		filter.Name = name
	}

	return filter
}
