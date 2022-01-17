package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/types/audits"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"testing"
)

func TestList(t *testing.T) {
	t.Parallel()

	tx := &transaction.RepositoryMock{}

	type args struct {
		ctx    context.Context
		filter account.Filter
	}

	type fields struct {
		tx   transaction.Transaction
		repo account.Repository
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    []Reference
		wantErr error
	}

	tests := []test{

		// Should return []Reference populated with information of accounts that satisfies the filter
		{
			name: "return array of accounts satisfying filter",
			fields: fields{tx: tx, repo: &account.RepositoryMock{ListResult: []account.Account{
				{
					ID:         1,
					ExternalID: id.ExternalFrom(accountID),
					Document:   "70830052062",
					Secret:     password.From("12345678"),
					Name:       "John Reis",
					Balance:    2500,
					Audit:      audits.Audit{},
				},
				{
					ID:         2,
					ExternalID: id.ExternalFrom(accountID),
					Document:   "24388516007",
					Secret:     password.From("12345678"),
					Name:       "Wagner Reis",
					Balance:    4500,
					Audit:      audits.Audit{},
				},
				{
					ID:         3,
					ExternalID: id.ExternalFrom(accountID),
					Document:   "05161964057",
					Secret:     password.From("12345678"),
					Name:       "Spencer Reis",
					Balance:    5000,
					Audit:      audits.Audit{},
				},
			}}},
			args: args{ctx: context.Background(), filter: account.Filter{Name: "Reis"}},
			want: []Reference{
				{ExternalID: id.ExternalFrom(accountID), Name: "John Reis"},
				{ExternalID: id.ExternalFrom(accountID), Name: "Wagner Reis"},
				{ExternalID: id.ExternalFrom(accountID), Name: "Spencer Reis"},
			},
		},

		// Should return []Reference empty due to no accounts satisfying filter
		{
			name:   "return empty array of accounts satisfying filter",
			fields: fields{tx: tx, repo: &account.RepositoryMock{ListResult: []account.Account{}}},
			args:   args{ctx: context.Background(), filter: account.Filter{Name: "Rise"}},
			want:   []Reference{},
		},

		// Should return ErrFetching on repository error
		{
			name:    "return ErrFetching on repository error",
			fields:  fields{tx: tx, repo: &account.RepositoryMock{Error: account.ErrFetching}},
			args:    args{ctx: context.Background(), filter: account.Filter{Name: "Reis"}},
			want:    []Reference{},
			wantErr: account.ErrFetching,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			u := New(test.fields.repo, test.fields.tx)
			got, err := u.List(test.args.ctx, test.args.filter)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
