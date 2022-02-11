package account

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	accountsDomain "stonehenge/app/core/entities/account"
	loggerDomain "stonehenge/app/core/types/logger"
	"stonehenge/app/gateway/api/account/schema"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/workspaces/account"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	t.Parallel()
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

	accounts := account.WorkspaceMock{
		ListResult: accountsDomain.GetFakeAccounts(),
	}
	logger := zerolog.New(os.Stdout)
	builder := rest.ResponseBuilder{
		Logger: logger,
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
			name:   "return 200 for successfully found account",
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
			ac := test.fields.accounts
			if ac == nil {
				ac = accounts
			}
			controller := NewController(ac, builder)

			var query string
			if test.args.params != nil {
				q := url.Values{}
				for k, v := range test.args.params {
					q.Add(k, v)
				}
				query = "?" + q.Encode()
			}
			req := httptest.NewRequest(http.MethodGet, "/accounts"+query, nil)

			router := chi.NewRouter()
			router.Method("GET", "/accounts", rest.Handler(controller.List))

			reqID := uuid.NewString()
			req = req.WithContext(context.WithValue(req.Context(), loggerDomain.RequestTracerContextKey, reqID))

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, test.wantCode, rec.Code)

			wantJSONBody, _ := json.Marshal(test.wantBody)
			assert.JSONEq(t, string(wantJSONBody), rec.Body.String())
		})
	}
}
