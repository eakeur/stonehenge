package account

import (
	"fmt"
	"net/http"
	"stonehenge/app/core/types/erring"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/account/schema"
	"stonehenge/app/gateway/api/rest"

	"github.com/go-chi/chi/v5"
)

// GetBalance gets the balance of the account specified
func (c *controller) GetBalance(r *http.Request) rest.Response {
	const operation = "Controller.Account.GetBalance"
	ctx := r.Context()
	accountID := id.ExternalFrom(chi.URLParam(r, "accountID"))
	balance, err := c.workspace.GetBalance(ctx, accountID)
	if err != nil {
		err = erring.Wrap(err, operation)
		return c.builder.BuildErrorResult(err).WithErrorLog(ctx)
	}
	return c.builder.
		BuildOKResult(schema.GetBalanceResponse{Balance: balance.Balance.ToStandardCurrency()}).
		WithSuccessLog(ctx, fmt.Sprintf("retrieved actual balance for account with ID %s", accountID))
}
