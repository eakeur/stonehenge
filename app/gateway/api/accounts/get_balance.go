package accounts

import (
	"net/http"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/accounts/schema"
	"stonehenge/app/gateway/api/rest"

	"github.com/go-chi/chi/v5"
)

// GetBalance gets the balance of the account specified
func (c *Controller) GetBalance(r *http.Request) rest.Response {
	const operation = "Controller.Account.GetBalance"
	ctx := r.Context()
	accountID := id.ExternalFrom(chi.URLParam(r, "id"))
	balance, err := c.workspace.GetBalance(ctx, accountID)
	if err != nil {
		c.logger.Error(ctx, operation, err.Error())
		return rest.BuildErrorResult(err)
	}
	c.logger.Trace(ctx, operation, "finished process successfully")
	return rest.BuildOKResult(schema.GetBalanceResponse{Balance: balance.Balance.ToStandardCurrency()})
}
