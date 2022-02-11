package transfer

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"stonehenge/app/core/entities/transfer"
	loggerDomain "stonehenge/app/core/types/logger"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/transfer/schema"
	transferWorker "stonehenge/app/worker/transfer"
	transferworkspace "stonehenge/app/workspaces/transfer"
	"testing"
)

func TestList(t *testing.T) {
	t.Parallel()
	const testWorkerTimeout = 5

	type fields struct {
		transfers transferworkspace.Workspace
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

	transfers := transferworkspace.WorkspaceMock{
		ListResult: transfer.GetFakeTransfers(),
	}

	logger := zerolog.New(os.Stdout)
	builder := rest.ResponseBuilder{
		Logger: logger,
	}

	var tests = []test{
		{
			name:     "return 200 for successfully found transfers",
			fields:   fields{},
			args:     args{},
			wantCode: http.StatusOK,
			wantBody: rest.Response{
				HTTPStatus: http.StatusOK,
				Content: func() []schema.ListTransferResponse {
					var res []schema.ListTransferResponse
					for _, v := range transfer.GetFakeTransfers() {
						res = append(res, schema.ListTransferResponse{
							ExternalID:    v.ExternalID.String(),
							Amount:        v.Amount.ToStandardCurrency(),
							EffectiveDate: v.EffectiveDate,
							OriginID:      v.Details.OriginExternalID.String(),
							DestinationID: v.Details.DestinationExternalID.String(),
						})
					}
					return res
				}(),
			},
		},
		{
			name:   "return 200 for successfully found transfers with filter",
			fields: fields{},
			args: args{
				params: map[string]string{
					"made_since": "2020-01-02",
					"made_until": "2020-02-02",
				},
			},
			wantCode: http.StatusOK,
			wantBody: rest.Response{
				HTTPStatus: http.StatusOK,
				Content: func() []schema.ListTransferResponse {
					var res []schema.ListTransferResponse
					for _, v := range transfer.GetFakeTransfers() {
						res = append(res, schema.ListTransferResponse{
							ExternalID:    v.ExternalID.String(),
							Amount:        v.Amount.ToStandardCurrency(),
							EffectiveDate: v.EffectiveDate,
							OriginID:      v.Details.OriginExternalID.String(),
							DestinationID: v.Details.DestinationExternalID.String(),
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
			tr := test.fields.transfers
			if tr == nil {
				tr = transfers
			}

			worker := transferWorker.NewWorker(testWorkerTimeout, tr, logger)
			controller := NewController(tr, worker, builder)

			var query string
			if test.args.params != nil {
				q := url.Values{}
				for k, v := range test.args.params {
					q.Add(k, v)
				}
				query = "?" + q.Encode()
			}
			req := httptest.NewRequest(http.MethodGet, "/transfers"+query, nil)

			router := chi.NewRouter()
			router.Method("GET", "/transfers", rest.Handler(controller.List))

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
