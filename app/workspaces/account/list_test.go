package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"testing"
)

func TestList(t *testing.T) {
	t.Parallel()

	accountID := id.ExternalFrom("d0052623-0695-4a3a-abf6-887f613dda8e")

	singleInstancePassword := password.From("D@V@C@O@")

	transactions := &transaction.RepositoryMock{}

	accesses := &access.RepositoryMock{
		GetAccessFromContextResult: access.Access{AccountID: accountID},
	}

	type args struct {
		ctx    context.Context
		filter account.Filter
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
		want    []account.Account
		wantErr error
	}

	tests := []test{

		{
			name: "return array of accounts",
			fields: fields{
				repo: &account.RepositoryMock{
					ListFunc: func(ctx context.Context, filter account.Filter) ([]account.Account, error) {
						list := make([]account.Account, 0)
						for _, a := range account.GetFakeAccounts() {
							a.Secret = singleInstancePassword
							list = append(list, a)
						}
						return list, nil
					},
				},
			},
			args: args{ctx: context.Background(), filter: account.Filter{}},
			want: func() []account.Account {
				list := make([]account.Account, 0)
				for _, a := range account.GetFakeAccounts() {
					a.Secret = singleInstancePassword
					list = append(list, a)
				}
				return list
			}(),
		},

		{
			name:   "return empty array",
			fields: fields{repo: &account.RepositoryMock{ListResult: []account.Account{}}},
			args:   args{ctx: context.Background()},
			want:   []account.Account{},
		},

		{
			name:    "return error on repository return error",
			fields:  fields{repo: &account.RepositoryMock{Error: account.ErrFetching}},
			args:    args{ctx: context.Background()},
			want:    []account.Account{},
			wantErr: account.ErrFetching,
		},

		{
			name:   "return ErrNoAccessInContext for not signed in actor",
			fields: fields{
				repo: &account.RepositoryMock{},
				access: access.RepositoryMock{Error: access.ErrNoAccessInContext},
			},
			args:   args{ctx: context.Background()},
			want:   []account.Account{},
			wantErr: access.ErrNoAccessInContext,
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
			got, err := u.List(test.args.ctx, test.args.filter)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
