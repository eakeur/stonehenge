package accounts

import (
	"net/http"
	"net/url"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/gateway/api/accounts/schema"
	"stonehenge/app/gateway/api/responses"
)

// List gets all accounts that satisfy the filter passed
func (c *Controller) List(r *http.Request) responses.Response {
	filters := filter(r.URL.Query())
	list, err := c.workspace.List(r.Context(), filters)
	if err != nil {
		return responses.BuildErrorResult(err)
	}
	res := make([]schema.ListResponse, len(list))
	for i := 0; i < len(list); i++ {
		res[i] = schema.ListResponse{
			AccountID: list[i].ExternalID.String(),
			OwnerName: list[i].Name,
		}
	}
	return responses.BuildOKResult(res)
}

func filter(values url.Values) account.Filter {
	f := account.Filter{}

	if name := values.Get("name"); name != "" {
		f.Name = name
	}

	return f
}
