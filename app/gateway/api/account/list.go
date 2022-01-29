package account

import (
	"fmt"
	"net/http"
	"net/url"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/erring"
	"stonehenge/app/gateway/api/account/schema"
	"stonehenge/app/gateway/api/rest"
)

// List gets all accounts that satisfy the filter passed
func (c *controller) List(r *http.Request) rest.Response {
	const operation = "Controller.Account.Create"
	ctx := r.Context()
	filters := filter(r.URL.Query())
	list, err := c.workspace.List(ctx, filters)
	if err != nil {
		err = erring.Wrap(err, operation)
		return c.builder.BuildErrorResult(err).WithErrorLog(ctx)
	}
	length := len(list)
	res := make([]schema.ListResponse, len(list))
	for i := 0; i < len(list); i++ {
		res[i] = schema.ListResponse{
			AccountID: list[i].ExternalID.String(),
			OwnerName: list[i].Name,
		}
	}
	return c.builder.BuildOKResult(res).
		AddHeaders("X-Total-Count", fmt.Sprint(length)).
		WithSuccessLog(ctx, fmt.Sprintf("listed accounts with %d results", length))
}

func filter(values url.Values) account.Filter {
	f := account.Filter{}

	if name := values.Get("name"); name != "" {
		f.Name = name
	}

	return f
}
