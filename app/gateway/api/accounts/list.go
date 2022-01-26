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
	const operation = "Controller.Account.Create"
	ctx := r.Context()
	filters := filter(r.URL.Query())
	list, err := c.workspace.List(ctx, filters)
	if err != nil {
		c.logger.Error(ctx, operation, err.Error())
		return rest.BuildErrorResult(err)
	}
	res := make([]schema.ListResponse, len(list))
	for i := 0; i < len(list); i++ {
		res[i] = schema.ListResponse{
			AccountID: list[i].ExternalID.String(),
			OwnerName: list[i].Name,
		}
	}
	c.logger.Trace(ctx, operation, "finished process successfully")
	return rest.BuildOKResult(res)
}

func filter(values url.Values) account.Filter {
	f := account.Filter{}

	if name := values.Get("name"); name != "" {
		f.Name = name
	}

	return f
}
