package accounts

import (
	"net/http"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/common"

	"github.com/go-chi/chi/v5"
)

// GetBalance gets the balance of the account specified
func (c *Controller) GetBalance(rw http.ResponseWriter, r *http.Request) {
	accountID := id.ID(chi.URLParam(r, "id"))
	if accountID != common.FetchContextUser(r.Context()) {
		// TODO Do not authorize other accounts to see one's balance
	}
	c.workspace.GetBalance(r.Context(), id.ID(accountID))
}
