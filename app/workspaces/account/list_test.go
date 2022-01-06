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
	type args struct {
		ctx    context.Context
		filter account.Filter
	}

	type fields struct {
		repo account.Repository
		tx   transaction.Transaction
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
				repo: &account.RepositoryMock{
					ListFunc: func(ctx context.Context, filter account.Filter) ([]account.Account, error) {
						return []account.Account{
							{
								ID:         1,
								ExternalID: id.New(),
								Document:   "70830052062",
								Secret:     password.From("12345678"),
								Name:       "John Reis",
								Balance:    2500,
								Audit:      audits.Audit{},
							},
							{
								ID:         2,
								ExternalID: id.New(),
								Document:   "24388516007",
								Secret:     password.From("12345678"),
								Name:       "Wagner Reis",
								Balance:    4500,
								Audit:      audits.Audit{},
							},
							{
								ID:         3,
								ExternalID: id.New(),
								Document:   "05161964057",
								Secret:     password.From("12345678"),
								Name:       "Spencer Reis",
								Balance:    5000,
								Audit:      audits.Audit{},
							},
						}, nil
					},
				},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx:    context.Background(),
				filter: account.Filter{Name: "Reis"},
			},
			want: []Reference{
				{
					Id:   1,
					Name: "John Reis",
				},
				{
					Id:   2,
					Name: "Wagner Reis",
				},
				{
					Id:   3,
					Name: "Spencer Reis",
				},
			},
		},
		{
			name: "should return unsuccessful list result due to no accounts found",
			fields: fields{
				repo: &account.RepositoryMock{
					ListFunc: func(ctx context.Context, filter account.Filter) ([]account.Account, error) {
						return []account.Account{}, account.ErrNotFound
					},
				},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx:    context.Background(),
				filter: account.Filter{Name: "Raising"},
			},
			want: nil,
			wantErr: account.ErrNotFound,
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
