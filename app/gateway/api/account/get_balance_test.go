package account

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	accountErrors "stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/id"
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

func TestGetBalance(t *testing.T) {
	t.Parallel()
	type fields struct {
		accounts account.Workspace
	}

	type args struct {
		id string
	}

	type test struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody rest.Response
	}

	accountIDMock := id.ExternalFrom("d0052623-0695-4a3a-abf6-887f613dda8e")
	accounts := account.WorkspaceMock{
		GetBalanceResult: account.GetBalanceResponse{Balance: 500},
	}
	logger := zerolog.New(os.Stdout)
	builder := rest.ResponseBuilder{
		Logger: logger,
	}

	var tests = []test{
		{
			name:     "return 200 for successfully found account",
			fields:   fields{},
			args:     args{id: accountIDMock.String()},
			wantCode: http.StatusOK,
			wantBody: rest.Response{
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
					Error: accountErrors.ErrNotFound,
				},
			},
			args:     args{id: accountIDMock.String()},
			wantCode: http.StatusNotFound,
			wantBody: rest.Response{
				HTTPStatus: http.StatusNotFound,
				Error:      rest.ErrAccountNotFound,
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

			req := httptest.NewRequest(http.MethodGet, "/accounts/"+test.args.id+"/balance", nil)

			router := chi.NewRouter()
			router.Method("GET", "/accounts/{accountID}/balance", rest.Handler(controller.GetBalance))

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
