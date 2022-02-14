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

func TestList(t *testing.T) {
	t.Parallel()

	accounts := account.WorkspaceMock{
		ListResult: accountDomain.GetFakeAccounts(),
	}

	type fields struct {
		accounts account.Workspace
	}

	type args struct {
		params map[string]string
	}

	type test struct {
		name     string
		fields   fields
		noAuth   bool
		args     args
		wantCode int
		wantBody rest.Response
	}

	var cases = []test{
		{
			name:     "return 200 for successfully found account",
			fields:   fields{},
			args:     args{},
			wantCode: http.StatusOK,
			wantBody: rest.Response{
				HTTPStatus: http.StatusOK,
				Content: func() []schema.AccountListResponse {
					var res []schema.AccountListResponse
					for _, v := range accountDomain.GetFakeAccounts() {
						res = append(res, schema.AccountListResponse{
							AccountID: v.ExternalID.String(),
							OwnerName: v.Name,
						})
					}
					return res
				}(),
			},
		},
		{
			name:   "return 200 for successfully found accounts with filter",
			fields: fields{},
			args: args{
				params: map[string]string{
					"name": "Igor",
				},
			},
			wantCode: http.StatusOK,
			wantBody: rest.Response{
				HTTPStatus: http.StatusOK,
				Content: func() []schema.AccountListResponse {
					var res []schema.AccountListResponse
					for _, v := range accountDomain.GetFakeAccounts() {
						res = append(res, schema.AccountListResponse{
							AccountID: v.ExternalID.String(),
							OwnerName: v.Name,
						})
					}
					return res
				}(),
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

			req := tests.CreateRequestWithParams(http.MethodGet, "/accounts", test.args.params)
			if !test.noAuth {
				req = tests.AuthenticateRequest(req, id.NewExternal())
			}
			rec := tests.Route{
				Method: http.MethodGet, Pattern: "/accounts",
				Handler: controller.List, RequiresAuth: true,
			}.ServeHTTP(req)

			assert.Equal(t, test.wantCode, rec.Code)
			wantJSONBody, _ := json.Marshal(test.wantBody)
			assert.JSONEq(t, string(wantJSONBody), rec.Body.String())
		})
	}
}
