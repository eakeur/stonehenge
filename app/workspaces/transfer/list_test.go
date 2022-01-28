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
	ctx := context.Background()

	originID := id.ExternalFrom("d0052623-0695-4a3a-abf6-887f613dda8e")
	//unknownID := id.ExternalFrom("17edb329-4b65-41ba-bb26-5060a1e157ab")

	accounts := &account.RepositoryMock{}

	transactions := &transaction.RepositoryMock{BeginResult: ctx}

	accesses := &access.RepositoryMock{
		GetAccessFromContextResult: access.Access{AccountID: originID},
	}

	type args struct {
		ctx    context.Context
		filter transfer.Filter
	}

	type fields struct {
		transactions transaction.Manager
		accounts     account.Repository
		transfers    transfer.Repository
		access       access.Manager
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    []transfer.Transfer
		wantErr error
	}

	tests := []test{

		{
			name:   "return array of transfers",
			fields: fields{transfers: &transfer.RepositoryMock{ListResult: transfer.GetFakeTransfers()}},
			args:   args{ctx: context.Background(), filter: transfer.Filter{OriginID: originID}},
			want:   transfer.GetFakeTransfers(),
		},

		{
			name:   "return empty array of transfers",
			fields: fields{transfers: &transfer.RepositoryMock{ListResult: []transfer.Transfer{}}},
			args:   args{ctx: context.Background(), filter: transfer.Filter{OriginID: originID}},
			want:   []transfer.Transfer{},
		},

		// Should return ErrFetching on repository error
		{
			name:    "return ErrFetching on repository error",
			fields:  fields{transfers: &transfer.RepositoryMock{Error: transfer.ErrFetching}},
			args:    args{ctx: context.Background(), filter: transfer.Filter{OriginID: originID}},
			want:    []transfer.Transfer{},
			wantErr: transfer.ErrFetching,
		},

		{
			name: "return ErrNoAccessInContext on no logged in user",
			fields: fields{
				access:    access.RepositoryMock{Error: access.ErrNoAccessInContext},
				transfers: &transfer.RepositoryMock{},
			},
			args:    args{ctx: context.Background(), filter: transfer.Filter{OriginID: originID}},
			want:    []transfer.Transfer{},
			wantErr: access.ErrNoAccessInContext,
		},

		{
			name: "return ErrCannotAccess on not authorized user",
			fields: fields{
				transfers: &transfer.RepositoryMock{},
			},
			args:    args{ctx: context.Background(), filter: transfer.Filter{}},
			want:    []transfer.Transfer{},
			wantErr: account.ErrCannotAccess,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tx := test.fields.transactions
			if tx == nil {
				tx = transactions
			}

			tk := test.fields.access
			if tk == nil {
				tk = accesses
			}

			ac := test.fields.accounts
			if ac == nil {
				ac = accounts
			}

			u := New(ac, test.fields.transfers, tx, tk)
			got, err := u.List(test.args.ctx, test.args.filter)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
