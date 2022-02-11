package account

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"stonehenge/app/core/entities/access"
	accountErrors "stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/id"
	loggerDomain "stonehenge/app/core/types/logger"
	"stonehenge/app/gateway/api/account/schema"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/workspaces/account"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type fields struct {
		accounts account.Workspace
	}

	type args struct {
		body schema.CreateAccountRequest
	}

	type test struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody rest.Response
	}

	accountIDMock := id.ExternalFrom("d0052623-0695-4a3a-abf6-887f613dda8e")
	accountTokenMock := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	createdAtMock := time.Now()

	accounts := account.WorkspaceMock{
		CreateResult: account.CreateOutput{
			AccountID: accountIDMock,
			CreatedAt: createdAtMock,
			Access: access.Access{
				AccountID: accountIDMock,
				Token:     accountTokenMock,
			},
		},
	}
	logger := zerolog.New(os.Stdout)
	builder := rest.ResponseBuilder{
		Logger: logger,
	}

	var tests = []test{
		{
			name:   "return 201 for successfully created account",
			fields: fields{},
			args: args{
				body: schema.CreateAccountRequest{
					Document: "49360206806",
					Secret:   "12345678",
					Name:     "Peanut Butter",
				},
			},
			wantCode: http.StatusCreated,
			wantBody: rest.Response{
				HTTPStatus: http.StatusCreated,
				Content: schema.CreateAccountResponse{
					AccountID: accountIDMock.String(),
					Token:     accountTokenMock,
				},
			},
		},
		{
			name: "return 400 for already existing account",
			fields: fields{
				accounts: account.WorkspaceMock{
					Error: accountErrors.ErrAlreadyExist,
				},
			},
			args: args{
				body: schema.CreateAccountRequest{
					Document: "49360206806",
					Secret:   "12345678",
					Name:     "Peanut Butter",
				},
			},
			wantCode: http.StatusBadRequest,
			wantBody: rest.Response{
				HTTPStatus: http.StatusBadRequest,
				Error:      rest.ErrAccountAlreadyExists,
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

			body, _ := json.Marshal(test.args.body)
			req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))

			router := chi.NewRouter()
			router.Method("POST", "/accounts", rest.Handler(controller.Create))

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
