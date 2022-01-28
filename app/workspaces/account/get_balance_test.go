package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/types/id"
	"testing"
)

func TestGetBalance(t *testing.T) {
	t.Parallel()

	accountID := id.ExternalFrom("d0052623-0695-4a3a-abf6-887f613dda8e")

	transactions := &transaction.RepositoryMock{}

	accesses := &access.RepositoryMock{
		GetAccessFromContextResult: access.Access{AccountID: accountID},
	}

	type args struct {
		ctx context.Context
		id  id.External
	}

	type fields struct {
		transactions transaction.Manager
		access       access.Manager
		repo         account.Repository
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    GetBalanceResponse
		wantErr error
	}

	tests := []test{

		// Should return the balance property of the account returned by the repository
		{
			name:   "return Balance property of account",
			fields: fields{repo: &account.RepositoryMock{GetBalanceResult: 5000}},
			args:   args{ctx: context.Background(), id: accountID},
			want:   GetBalanceResponse{Balance: 5000},
		},

		{
			name:   "return ErrNoAccessInContext for not signed in actor",
			fields: fields{
				repo: &account.RepositoryMock{},
				access: access.RepositoryMock{Error: access.ErrNoAccessInContext},
			},
			args:   args{ctx: context.Background(), id: accountID},
			wantErr: access.ErrNoAccessInContext,
		},

		{
			name:   "return ErrCannotAccess for not authorized actor",
			fields: fields{
				repo: &account.RepositoryMock{},
				access: access.RepositoryMock{
					GetAccessFromContextResult: access.Access{AccountID: id.NewExternal()},
				},
			},
			args:   args{ctx: context.Background(), id: accountID},
			wantErr: account.ErrCannotAccess,
		},

		// Should return ErrNotFound because no account with the ID specified was found
		{
			name:    "return ErrNotFound for nonexistent account",
			fields:  fields{repo: &account.RepositoryMock{Error: account.ErrNotFound}},
			args:    args{ctx: context.Background(), id: accountID},
			wantErr: account.ErrNotFound,
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

			u := New(test.fields.repo, tx, tk)
			got, err := u.GetBalance(test.args.ctx, test.args.id)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
