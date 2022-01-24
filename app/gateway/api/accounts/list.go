package accounts

import (
	"net/http"
	"net/url"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/gateway/api/accounts/schema"
	"stonehenge/app/gateway/api/rest"
)

// List gets all accounts that satisfy the filter passed
func (c *Controller) List(r *http.Request) rest.Response {
	filters := filter(r.URL.Query())
	list, err := c.workspace.List(r.Context(), filters)
	if err != nil {
		return rest.BuildErrorResult(err)
	}
	res := make([]schema.ListResponse, len(list))
	for i := 0; i < len(list); i++ {
		res[i] = schema.ListResponse{
			AccountID: list[i].ExternalID.String(),
			OwnerName: list[i].Name,
		}
	}
	return rest.BuildOKResult(res)
}

func filter(values url.Values) account.Filter {
	f := account.Filter{}

	if name := values.Get("name"); name != "" {
		f.Name = name
	}

	return f
}
