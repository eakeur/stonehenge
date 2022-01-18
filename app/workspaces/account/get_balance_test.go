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

	tx := &transaction.RepositoryMock{}

	tk := &access.FactoryMock{
		GetAccessFromContextResult: access.Access{AccountID: id.ExternalFrom(accountID)},
	}

	type args struct {
		ctx context.Context
		id  id.External
	}

	type fields struct {
		tx   transaction.Transaction
		tk   access.Factory
		repo account.Repository
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
			fields: fields{tx: tx, tk: tk, repo: &account.RepositoryMock{GetBalanceResult: 5000}},
			args:   args{ctx: context.Background(), id: id.ExternalFrom(accountID)},
			want:   GetBalanceResponse{Balance: 5000},
		},

		// Should return ErrNotFound because no account with the ID specified was found
		{
			name:    "return ErrNotFound for nonexistent account",
			fields:  fields{tx: tx, tk: tk, repo: &account.RepositoryMock{Error: account.ErrNotFound}},
			args:    args{ctx: context.Background(), id: id.ExternalFrom(accountID)},
			wantErr: account.ErrNotFound,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			u := New(test.fields.repo, test.fields.tx, test.fields.tk)
			got, err := u.GetBalance(test.args.ctx, test.args.id)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
