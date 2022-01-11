package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/types/audits"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthentication(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx   context.Context
		input AuthenticationRequest
	}

	type fields struct {
		repo account.Repository
		tx   transaction.Transaction
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    id.ExternalID
		wantErr error
	}
	var accountID = id.New()

	var tests = []test{
		{
			name: "should return successful authentication",
			fields: fields{
				repo: &account.RepositoryMock{
					GetWithCPFFunc: func(ctx context.Context, document document.Document) (account.Account, error) {
						return account.Account{
							ID:         1,
							ExternalID: accountID,
							Document:   "97662062015",
							Secret:     password.From("12345678"),
							Name:       "Lina Pereira",
							Balance:    5000,
							Audit:      audits.Audit{},
						}, nil
					},
				},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				input: AuthenticationRequest{
					Document: "97662062015",
					Secret:   "12345678",
				},
			},
			want: accountID,
		},
		{
			name: "should return unsuccessful authentication due to wrong password",
			fields: fields{
				repo: &account.RepositoryMock{
					GetWithCPFFunc: func(ctx context.Context, document document.Document) (account.Account, error) {
						return account.Account{
							ID:         1,
							ExternalID: accountID,
							Document:   "97662062015",
							Secret:     password.From("12345678"),
							Name:       "Lina Pereira",
							Balance:    5000,
							Audit:      audits.Audit{},
						}, nil
					},
				},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				input: AuthenticationRequest{
					Document: "97662062015",
					Secret:   "12345677",
				},
			},
			want:    id.ZeroValue,
			wantErr: password.ErrWrongPassword,
		},
		{
			name: "should return unsuccessful authentication due invalid document",
			fields: fields{
				repo: &account.RepositoryMock{},
				tx:   &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				input: AuthenticationRequest{
					Document: "97662062",
					Secret:   "12345677",
				},
			},
			want:    id.ZeroValue,
			wantErr: document.ErrInvalidDocument,
		},
		{
			name: "should return unsuccessful due to account not found error",
			fields: fields{
				repo: &account.RepositoryMock{
					GetWithCPFFunc: func(ctx context.Context, document document.Document) (account.Account, error) {
						return account.Account{}, account.ErrNotFound
					},
				},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				input: AuthenticationRequest{
					Document: "97662062015",
					Secret:   "12345678",
				},
			},
			want:    id.ZeroValue,
			wantErr: account.ErrNotFound,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			u := New(test.fields.repo, test.fields.tx)
			got, err := u.Authenticate(test.args.ctx, test.args.input)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
