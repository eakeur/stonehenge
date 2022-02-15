package transfer

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/tests"
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
		noAuth   bool
		wantBody rest.Response
	}

	transfers := transferworkspace.WorkspaceMock{
		ListResult: transfer.GetFakeTransfers(),
	}

	var cases = []test{
		{
			name:     "return 200 for successfully found transfers",
			fields:   fields{},
			args:     args{},
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

			req := tests.CreateRequestWithParams(http.MethodGet, "/transfers", test.args.params)
			rec := tests.Route{Method: http.MethodGet, Pattern: "/transfers", Handler: controller.List}.
				ServeHTTP(req)

			assert.Equal(t, test.wantBody.HTTPStatus, rec.Code)
			wantJSONBody, _ := json.Marshal(test.wantBody)
			assert.JSONEq(t, string(wantJSONBody), rec.Body.String())
		})
	}
}
