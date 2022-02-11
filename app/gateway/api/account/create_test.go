package account

import (
	"encoding/json"
	"net/http"
	"stonehenge/app/core/entities/access"
	accountDomain "stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/account/schema"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/tests"
	"stonehenge/app/workspaces/account"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	accountIDMock := id.ExternalFrom("d0052623-0695-4a3a-abf6-887f613dda8e")
	createdAtMock := time.Now()

	accounts := account.WorkspaceMock{
		CreateResult: account.CreateOutput{
			AccountID: accountIDMock,
			CreatedAt: createdAtMock,
			Access: access.Access{
				AccountID: accountIDMock,
				Token:     tests.JWTTokenMock,
			},
		},
	}

	type fields struct {
		accounts account.Workspace
	}

	type args struct {
		body interface{}
	}

	type test struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody rest.Response
	}

	var cases = []test{
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
					Token:     tests.JWTTokenMock,
				},
			},
		},
		{
			name: "return 400 for already existing account",
			fields: fields{
				accounts: account.WorkspaceMock{
					Error: accountDomain.ErrAlreadyExist,
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

	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			controller := NewController(
				tests.EvaluateDep(test.fields.accounts, accounts).(account.Workspace),
				tests.GetResponseBuilder(),
			)

			req := tests.CreateRequestWithBody(http.MethodPost, "/accounts", test.args.body)
			rec := tests.Route{Method: http.MethodPost, Pattern: "/accounts", Handler: controller.Create}.
				ServeHTTP(req)

			assert.Equal(t, test.wantCode, rec.Code)
			wantJSONBody, _ := json.Marshal(test.wantBody)
			assert.JSONEq(t, string(wantJSONBody), rec.Body.String())
		})
	}
}
