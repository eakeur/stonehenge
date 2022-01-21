package accounts

import (
	"net/http"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/accounts/schema"
	"stonehenge/app/gateway/api/responses"

	"github.com/go-chi/chi/v5"
)

// GetBalance gets the balance of the account specified
func (c *Controller) GetBalance(r *http.Request) responses.Response {
	accountID := id.ExternalFrom(chi.URLParam(r, "id"))
	balance, err := c.workspace.GetBalance(r.Context(), accountID)
	if err != nil {
		return responses.BuildErrorResult(err)
	}
	return responses.BuildOKResult(schema.GetBalanceResponse{Balance: float64(balance.Balance)})
}
