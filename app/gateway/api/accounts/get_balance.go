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
	accountID := id.ExternalFrom(chi.URLParam(r, "id"))
	balance, err := c.workspace.GetBalance(r.Context(), accountID)
	if err != nil {
		return rest.BuildErrorResult(err)
	}
	return rest.BuildOKResult(schema.GetBalanceResponse{Balance: float64(balance.Balance)})
}
