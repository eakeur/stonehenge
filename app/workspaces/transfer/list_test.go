package transfer

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/audits"
	"stonehenge/app/core/types/id"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		filter transfer.Filter
	}

	type fields struct {
		ac account.Repository
		tr transfer.Repository
		tx transaction.Transaction
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    []Reference
		wantErr error
	}

	var tests = []test{
		{
			name: "should return successful list with results",
			fields: fields{
				ac: &account.RepositoryMock{},
				tr: &transfer.RepositoryMock{
					ListFunc: func(c context.Context, f transfer.Filter) ([]transfer.Transfer, error) {
						return []transfer.Transfer{
							{
								ID:            1,
								ExternalID:    id.ZeroValue,
								OriginID:      1,
								DestinationID: 2,
								Amount:        5,
								EffectiveDate: time.Time{},
								Audit:         audits.Audit{},
							},
						}, nil
					},
				},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx:    context.Background(),
				filter: transfer.Filter{},
			},
			want: []Reference{
				{
					Id:            1,
					OriginId:      1,
					DestinationId: 2,
					Amount:        5,
					EffectiveDate: time.Time{},
				},
			},
		},
		{
			name: "should return unsuccessful list result due to no transfers found",
			fields: fields{
				ac: &account.RepositoryMock{},
				tr: &transfer.RepositoryMock{
					ListFunc: func(c context.Context, f transfer.Filter) ([]transfer.Transfer, error) {
						return []transfer.Transfer{}, transfer.ErrNotFound
					},
				},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx:    context.Background(),
				filter: transfer.Filter{},
			},
			want:    nil,
			wantErr: transfer.ErrNotFound,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			u := New(test.fields.ac, test.fields.tr, test.fields.tx)
			got, err := u.List(test.args.ctx, test.args.filter)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
