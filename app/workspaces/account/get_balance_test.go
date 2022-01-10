package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"testing"
)

func TestGetBalance(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
		id  id.ExternalID
	}

	type fields struct {
		repo account.Repository
		tx   transaction.Transaction
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    GetBalanceResponse
		wantErr error
	}

	var tests = []test{
		{
			name: "should return successful balance result",
			fields: fields{
				repo: &account.RepositoryMock{
					GetBalanceFunc: func(ctx context.Context, id id.ExternalID) (currency.Currency, error) {
						return 5000, nil
					},
				},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				id:  id.New(),
			},
			want: GetBalanceResponse{Balance: 5000},
		},
		{
			name: "should return unsuccessful balance result due to not found error",
			fields: fields{
				repo: &account.RepositoryMock{
					GetBalanceFunc: func(ctx context.Context, id id.ExternalID) (currency.Currency, error) {
						return 0, account.ErrNotFound
					},
				},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				id:  id.New(),
			},
			want:    GetBalanceResponse{},
			wantErr: account.ErrNotFound,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			u := New(test.fields.repo, test.fields.tx)
			got, err := u.GetBalance(test.args.ctx, test.args.id)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
