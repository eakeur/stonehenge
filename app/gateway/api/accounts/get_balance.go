package accounts

import (
	"errors"
	"net/http"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/common"
	"stonehenge/app/gateway/api/responses"

	"github.com/go-chi/chi/v5"
)

// GetBalance gets the balance of the account specified
func (c *Controller) GetBalance(rw http.ResponseWriter, r *http.Request) {
	accountID := id.ID(chi.URLParam(r, "id"))
	ctx := r.Context()
	if accountID != common.FetchContextUser(ctx) {
		responses.WriteErrorResponse(rw, http.StatusForbidden, errors.New("you do not have access to this resource"))
		return
	}
	balance, err := c.workspace.GetBalance(ctx, accountID)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusBadRequest, err)
		return
	}

	err = responses.WriteSuccessfulJSON(rw, http.StatusOK, balance)
	if err != nil {
		responses.WriteErrorResponse(rw, http.StatusInternalServerError, err)
	}
}
