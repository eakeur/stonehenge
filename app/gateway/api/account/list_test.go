package account

import (
	"encoding/json"
	"net/http"
	accountsDomain "stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/account/schema"
	"stonehenge/app/gateway/api/rest"
	testutils "stonehenge/app/test_utils"
	"stonehenge/app/workspaces/account"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	t.Parallel()

	accounts := account.WorkspaceMock{
		ListResult: accountsDomain.GetFakeAccounts(),
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
		args     args
		wantCode int
		wantBody rest.Response
	}

	var tests = []test{
		{
			name:     "return 200 for successfully found account",
			fields:   fields{},
			args:     args{},
			wantCode: http.StatusOK,
			wantBody: rest.Response{
				HTTPStatus: http.StatusOK,
				Content: func() []schema.AccountListResponse {
					var res []schema.AccountListResponse
					for _, v := range accountsDomain.GetFakeAccounts() {
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
					for _, v := range accountsDomain.GetFakeAccounts() {
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

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			controller := NewController(
				testutils.EvaluateDep(test.fields.accounts, accounts).(account.Workspace),
				testutils.GetResponseBuilder(),
			)

			req := testutils.CreateRequestWithParams(http.MethodGet, "/accounts", test.args.params)
			req = testutils.AuthenticateRequest(req, id.NewExternal())
			rec := testutils.Route{
				Method: http.MethodGet, Pattern: "/accounts",
				Handler: controller.List, RequiresAuth: true,
			}.ServeHTTP(req)

			assert.Equal(t, test.wantCode, rec.Code)
			wantJSONBody, _ := json.Marshal(test.wantBody)
			assert.JSONEq(t, string(wantJSONBody), rec.Body.String())
		})
	}
}
