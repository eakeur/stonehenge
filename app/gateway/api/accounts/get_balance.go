package accounts

import (
	"net/http"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/responses"

	"github.com/go-chi/chi/v5"
)

// GetBalance gets the balance of the account specified
func (c *Controller) GetBalance(rw http.ResponseWriter, r *http.Request) {
	accountID := id.ExternalFrom(chi.URLParam(r, "id"))
	balance, err := c.workspace.GetBalance(r.Context(), accountID)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusBadRequest, err)
		return
	}

	err = responses.WriteSuccessfulJSON(rw, http.StatusOK, balance)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusInternalServerError, err)
	}
}
