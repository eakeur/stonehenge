package transfer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/id"
	"testing"
)

func TestList(t *testing.T) {
	t.Parallel()
	tx := &transaction.RepositoryMock{}
	ac := &account.RepositoryMock{}
	tk := &access.RepositoryMock{
		CreateResult: access.Access{
			AccountID: id.ExternalFrom(originID),
		},
	}

	type args struct {
		ctx    context.Context
		filter transfer.Filter
	}

	type fields struct {
		tx transaction.Transaction
		ac account.Repository
		tr transfer.Repository
		tk access.Manager
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    []Reference
		wantErr error
	}

	tests := []test{
		// Should return []Reference populated with information of transfers that satisfies the filter
		{
			name: "return array of transfers satisfying filter",
			fields: fields{tx: tx, tk: tk, ac: ac, tr: &transfer.RepositoryMock{ListResult: []transfer.Transfer{
				{
					ID:            1,
					ExternalID:    id.ExternalFrom(transferID),
					OriginID:      1,
					DestinationID: 2,
					Amount:        2500,
				},
				{
					ID:            2,
					ExternalID:    id.ExternalFrom(transferID),
					OriginID:      1,
					DestinationID: 3,
					Amount:        440,
				},
				{
					ID:            3,
					ExternalID:    id.ExternalFrom(transferID),
					OriginID:      1,
					DestinationID: 4,
					Amount:        50000,
				},
				{
					ID:            4,
					ExternalID:    id.ExternalFrom(transferID),
					OriginID:      1,
					DestinationID: 5,
					Amount:        2660,
				},
			}}},
			args: args{ctx: context.Background(), filter: transfer.Filter{OriginID: id.ExternalFrom(originID)}},
			want: []Reference{
				{
					ExternalID:    id.ExternalFrom(transferID),
					OriginID:      1,
					DestinationID: 2,
					Amount:        2500,
				},
				{
					ExternalID:    id.ExternalFrom(transferID),
					OriginID:      1,
					DestinationID: 3,
					Amount:        440,
				},
				{
					ExternalID:    id.ExternalFrom(transferID),
					OriginID:      1,
					DestinationID: 4,
					Amount:        50000,
				},
				{
					ExternalID:    id.ExternalFrom(transferID),
					OriginID:      1,
					DestinationID: 5,
					Amount:        2660,
				},
			},
		},

		// Should return []Reference empty due to no transfers satisfying filter
		{
			name:   "return empty array of transfers satisfying filter",
			fields: fields{tx: tx, tk: tk, ac: ac, tr: &transfer.RepositoryMock{ListResult: []transfer.Transfer{}}},
			args:   args{ctx: context.Background(), filter: transfer.Filter{OriginID: id.ExternalFrom(unknownID)}},
			want:   []Reference{},
		},

		// Should return ErrFetching on repository error
		{
			name:    "return ErrFetching on repository error",
			fields:  fields{tx: tx, tk: tk, ac: ac, tr: &transfer.RepositoryMock{Error: transfer.ErrFetching}},
			args:    args{ctx: context.Background(), filter: transfer.Filter{OriginID: id.ExternalFrom(unknownID)}},
			want:    []Reference{},
			wantErr: transfer.ErrFetching,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			u := New(test.fields.ac, test.fields.tr, test.fields.tx, test.fields.tk)
			got, err := u.List(test.args.ctx, test.args.filter)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
