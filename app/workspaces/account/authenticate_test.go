package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"testing"
)

const (
	accountID = "d0052623-0695-4a3a-abf6-887f613dda8e"
)

func TestAuthentication(t *testing.T) {
	t.Parallel()

	tx := &transaction.RepositoryMock{}

	type args struct {
		ctx   context.Context
		input AuthenticationRequest
	}

	type fields struct {
		tx   transaction.Transaction
		repo account.Repository
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    id.External
		wantErr error
	}

	var tests = []test{

		// Should return External of the successfully authenticated account
		{
			name: "return External of authenticated account",
			fields: fields{tx: tx, repo: &account.RepositoryMock{
				GetWithCPFResult: account.Account{
					ID:         1,
					ExternalID: id.ExternalFrom(accountID),
					Document:   "97662062015",
					Secret:     password.From("D@V@C@O@"),
					Name:       "Lina Pereira",
					Balance:    5000,
				},
			}},
			args: args{ctx: context.Background(), input: AuthenticationRequest{
				Document: "97662062015",
				Secret:   "D@V@C@O@",
			}},
			want: id.ExternalFrom(accountID),
		},

		// Should return ErrWrongPassword authenticating with unmatching password
		{
			name: "return ErrWrongPassword on unmatching password",
			fields: fields{tx: tx, repo: &account.RepositoryMock{
				GetWithCPFResult: account.Account{
					ID:         1,
					ExternalID: id.ExternalFrom(accountID),
					Document:   "97662062015",
					Secret:     password.From("D@V@C@O@"),
					Name:       "Lina Pereira",
					Balance:    5000,
				},
			}},
			args: args{ctx: context.Background(), input: AuthenticationRequest{
				Document: "97662062015",
				Secret:   "D@V@C@A@",
			}},
			want:    id.Zero,
			wantErr: password.ErrWrongPassword,
		},

		// Should return ErrInvalidDocument authenticating with corrupted CPF
		{
			name:   "return ErrInvalidDocument on corrupted CPF",
			fields: fields{tx: tx, repo: &account.RepositoryMock{}},
			args: args{ctx: context.Background(), input: AuthenticationRequest{
				Document: "9766206201",
				Secret:   "D@V@C@O@",
			}},
			want:    id.Zero,
			wantErr: document.ErrInvalidDocument,
		},

		// Should return ErrNotFound when authenticating a nonexistent account
		{
			name:   "return ErrNotFound authenticating nonexistent account",
			fields: fields{tx: tx, repo: &account.RepositoryMock{Error: account.ErrNotFound}},
			args: args{ctx: context.Background(), input: AuthenticationRequest{
				Document: "97662062015",
				Secret:   "D@V@C@O@",
			}},
			want:    id.Zero,
			wantErr: account.ErrNotFound,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			u := New(test.fields.repo, test.fields.tx)
			got, err := u.Authenticate(test.args.ctx, test.args.input)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
