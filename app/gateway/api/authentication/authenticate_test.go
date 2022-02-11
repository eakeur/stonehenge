package authentication

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/types/id"
	loggerDomain "stonehenge/app/core/types/logger"
	"stonehenge/app/core/types/password"
	"stonehenge/app/gateway/api/authentication/schema"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/workspaces/authentication"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type fields struct {
		auth authentication.Workspace
	}

	type args struct {
		body schema.AuthenticationRequest
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

	accounts := authentication.WorkspaceMock{
		AuthenticateResult: access.Access{
			AccountID:   accountIDMock,
			AccountName: "Igor Reis",
			Token:       accountTokenMock,
		},
	}

	logger := zerolog.New(os.Stdout)
	builder := rest.ResponseBuilder{
		Logger: logger,
	}

	var tests = []test{
		{
			name:   "return 200 for successfully authenticated account",
			fields: fields{},
			args: args{
				body: schema.AuthenticationRequest{
					Document: "49366655587",
					Secret:   "12345678",
				},
			},
			wantCode: http.StatusOK,
			wantBody: rest.Response{
				HTTPStatus: http.StatusOK,
				Content: schema.AuthenticationResponse{
					Token: accountTokenMock,
				},
			},
		},
		{
			name: "return 400 for bad password",
			fields: fields{
				auth: authentication.WorkspaceMock{Error: password.ErrWrongPassword},
			},
			args: args{
				body: schema.AuthenticationRequest{
					Document: "49366655587",
					Secret:   "123456787",
				},
			},
			wantCode: http.StatusUnauthorized,
			wantBody: rest.Response{
				HTTPStatus: http.StatusUnauthorized,
				Error:      rest.ErrPasswordInvalid,
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ac := test.fields.auth
			if ac == nil {
				ac = accounts
			}
			controller := NewController(ac, builder)

			body, _ := json.Marshal(test.args.body)
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))

			router := chi.NewRouter()
			router.Method("POST", "/login", rest.Handler(controller.Authenticate))

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
