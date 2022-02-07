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

// GetBalance godoc
// @Summary      Gets account balance
// @Description  Gets the balance of the account specified, if it's the logged in account
// @Tags         Accounts
// @Param       accountID   path string  true  "Account ID"
// @Produce      json
// @Success      200  {object}  schema.GetBalanceResponse
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Security     AuthKey
// @Router       /api/v1/accounts/{accountID}/balance [get]
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
