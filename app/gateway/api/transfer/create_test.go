package transfer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/id"
	loggerDomain "stonehenge/app/core/types/logger"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/transfer/schema"
	transferWorker "stonehenge/app/worker/transfer"
	transferworkspace "stonehenge/app/workspaces/transfer"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

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
		wantCode int
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

	logger := zerolog.New(os.Stdout)
	builder := rest.ResponseBuilder{
		Logger: logger,
	}

	var (
		tests = []test{
			{
				name:   "return 200 for successfully created transfer",
				fields: fields{},
				args: args{
					body: schema.CreateTransferRequest{
						DestinationID: destinationIDMock.String(),
						Amount:        500,
					},
				},
				wantCode: http.StatusCreated,
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
				wantCode: http.StatusBadRequest,
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
				wantCode: http.StatusBadRequest,
				wantBody: rest.Response{
					HTTPStatus: http.StatusBadRequest,
					Error:      rest.ErrTransferSameAccount,
				},
			},
		}
	)

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			tr := test.fields.transfers
			if tr == nil {
				tr = transfers
			}

			worker := transferWorker.NewWorker(testWorkerTimeout, tr, logger)
			go worker.Run()
			defer worker.Close()

			controller := NewController(tr, worker, builder)

			body, _ := json.Marshal(test.args.body)
			req := httptest.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(body))

			router := chi.NewRouter()
			router.Method("POST", "/transfers", rest.Handler(controller.Create))

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
