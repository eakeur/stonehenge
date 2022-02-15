package transfer

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/tests"
	"stonehenge/app/gateway/api/transfer/schema"
	transferWorker "stonehenge/app/worker/transfer"
	transferworkspace "stonehenge/app/workspaces/transfer"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	const testWorkerTimeout = 5

	type fields struct {
		transfers transferworkspace.Workspace
	}

	type args struct {
		body schema.CreateTransferRequest
	}

	type test struct {
		name     string
		fields   fields
		args     args
		noAuth   bool
		wantBody rest.Response
	}

	transferIDMock := id.ExternalFrom("d0052623-0695-4a3a-abf6-887f613dda8e")
	destinationIDMock := id.ExternalFrom("d0052623-0695-4a3a-abf6-887f613ddeee")
	createdAtMock := time.Now()

	transfers := transferworkspace.WorkspaceMock{
		CreateResult: transferworkspace.CreateOutput{
			RemainingBalance: 500,
			CreatedAt:        createdAtMock,
			TransferID:       transferIDMock,
		},
	}

	var (
		cases = []test{
			{
				name:   "return 200 for successfully created transfer",
				fields: fields{},
				args: args{
					body: schema.CreateTransferRequest{
						DestinationID: destinationIDMock.String(),
						Amount:        500,
					},
				},
				wantBody: rest.Response{
					HTTPStatus: http.StatusCreated,
					Content: schema.CreateTransferResponse{
						RemainingBalance: 5,
						TransferID:       transferIDMock.String(),
					},
				},
			},
			{
				name: "return 400 for account with not enough money",
				fields: fields{
					transfers: transferworkspace.WorkspaceMock{
						Error: account.ErrNoMoney,
					},
				},
				args: args{
					body: schema.CreateTransferRequest{
						DestinationID: destinationIDMock.String(),
						Amount:        500,
					},
				},
				wantBody: rest.Response{
					HTTPStatus: http.StatusBadRequest,
					Error:      rest.ErrAccountNoMoney,
				},
			},
			{
				name: "return 400 for same account account transfer",
				fields: fields{
					transfers: transferworkspace.WorkspaceMock{
						Error: transfer.ErrSameAccount,
					},
				},
				args: args{
					body: schema.CreateTransferRequest{
						DestinationID: destinationIDMock.String(),
						Amount:        500,
					},
				},
				wantBody: rest.Response{
					HTTPStatus: http.StatusBadRequest,
					Error:      rest.ErrTransferSameAccount,
				},
			},
		}
	)

	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			worker := transferWorker.NewWorker(
				testWorkerTimeout,
				tests.EvaluateDep(test.fields.transfers, transfers).(transferworkspace.Workspace),
				tests.GetResponseBuilder().Logger,
			)
			go worker.Run()
			defer worker.Close()

			controller := NewController(
				tests.EvaluateDep(test.fields.transfers, transfers).(transferworkspace.Workspace),
				worker,
				tests.GetResponseBuilder(),
			)

			req := tests.CreateRequestWithBody(http.MethodPost, "/transfers", test.args.body)
			rec := tests.Route{Method: http.MethodPost, Pattern: "/transfers", Handler: controller.Create}.
				ServeHTTP(req)

			assert.Equal(t, test.wantBody.HTTPStatus, rec.Code)
			wantJSONBody, _ := json.Marshal(test.wantBody)
			assert.JSONEq(t, string(wantJSONBody), rec.Body.String())
		})
	}
}
