package account

import (
	"encoding/json"
	"net/http"
	accountDomain "stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/account/schema"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/tests"
	"stonehenge/app/workspaces/account"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	t.Parallel()

	accountIDMock := id.ExternalFrom("d0052623-0695-4a3a-abf6-887f613dda8e")
	accounts := account.WorkspaceMock{
		GetBalanceResult: account.GetBalanceResponse{Balance: 500},
	}

	type fields struct {
		accounts account.Workspace
	}

	type args struct {
		id string
	}

	type test struct {
		name   string
		fields fields
		noAuth bool
		args   args
		want   rest.Response
	}

	var cases = []test{
		{
			name: "return 200 for successfully found account",
			args: args{id: accountIDMock.String()},
			want: rest.Response{
				HTTPStatus: http.StatusOK,
				Content: schema.GetBalanceResponse{
					Balance: 5.00,
				},
			},
		},
		{
			name: "return 404 for not found account",
			fields: fields{
				accounts: account.WorkspaceMock{
					Error: accountDomain.ErrNotFound,
				},
			},
			args: args{id: accountIDMock.String()},
			want: rest.Response{
				HTTPStatus: http.StatusNotFound,
				Error:      rest.ErrAccountNotFound,
			},
		},
		{
			name:   "return 401 for account not authenticated",
			args:   args{id: accountIDMock.String()},
			noAuth: true,
			want: rest.Response{
				HTTPStatus: http.StatusUnauthorized,
				Error:      rest.ErrAccessNonexistent,
			},
		},
	}

	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			controller := NewController(
				tests.EvaluateDep(test.fields.accounts, accounts).(account.Workspace),
				tests.GetResponseBuilder(),
			)

			req := tests.CreateRequestWithParams(http.MethodGet, "/accounts/"+test.args.id+"/balance", nil)
			if !test.noAuth {
				req = tests.AuthenticateRequest(req, id.NewExternal())
			}
			rec := tests.Route{
				Method: http.MethodGet, Pattern: "/accounts/{accountID}/balance",
				Handler: controller.GetBalance, RequiresAuth: true,
			}.ServeHTTP(req)

			assert.Equal(t, test.want.HTTPStatus, rec.Code)
			wantJSONBody, _ := json.Marshal(test.want)
			assert.JSONEq(t, string(wantJSONBody), rec.Body.String())
		})
	}
}
