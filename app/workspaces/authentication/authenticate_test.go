package authentication

import (
	"context"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthentication(t *testing.T) {
	t.Parallel()

	accountID := id.ExternalFrom("d0052623-0695-4a3a-abf6-887f613dda8e")

	accesses := &access.RepositoryMock{
		CreateResult: access.Access{AccountID: accountID},
	}

	type args struct {
		ctx   context.Context
		input AuthenticationRequest
	}

	type fields struct {
		access access.Manager
		repo   account.Repository
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    access.Access
		wantErr error
	}

	var tests = []test{

		{
			name: "return External of authenticated account",
			fields: fields{repo: &account.RepositoryMock{
				GetWithCPFResult: func() account.Account {
					ac := account.GetFakeAccount()
					ac.ExternalID = accountID
					return ac
				}(),
			}},
			args: args{ctx: context.Background(), input: AuthenticationRequest{
				Document: "24788516002",
				Secret:   "12345678",
			}},
			want: access.Access{AccountID: accountID},
		},

		{
			name: "return ErrWrongPassword on unmatching password",
			fields: fields{repo: &account.RepositoryMock{
				GetWithCPFResult: account.GetFakeAccount(),
			}},
			args: args{ctx: context.Background(), input: AuthenticationRequest{
				Document: "24788516002",
				Secret:   "12345677",
			}},
			wantErr: password.ErrWrongPassword,
		},

		{
			name:   "return ErrInvalidDocument on corrupted CPF",
			fields: fields{repo: &account.RepositoryMock{}},
			args: args{ctx: context.Background(), input: AuthenticationRequest{
				Document: "2478851600",
				Secret:   "12345678",
			}},
			wantErr: document.ErrInvalidDocument,
		},

		{
			name:   "return ErrNotFound authenticating nonexistent account",
			fields: fields{repo: &account.RepositoryMock{Error: account.ErrNotFound}},
			args: args{ctx: context.Background(), input: AuthenticationRequest{
				Document: "24788516002",
				Secret:   "12345678",
			}},
			wantErr: account.ErrNotFound,
		},

		{
			name: "return ErrTokenFailedCreation for unsuccessful token encryption",
			fields: fields{
				access: access.RepositoryMock{
					Error: access.ErrTokenFailedCreation,
				},
				repo: &account.RepositoryMock{
					GetWithCPFResult: func() account.Account {
						ac := account.GetFakeAccount()
						ac.ExternalID = accountID
						return ac
					}(),
				},
			},
			args: args{ctx: context.Background(), input: AuthenticationRequest{
				Document: "24788516002",
				Secret:   "12345678",
			}},
			wantErr: access.ErrTokenFailedCreation,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tk := test.fields.access
			if tk == nil {
				tk = accesses
			}

			u := New(test.fields.repo, tk)
			got, err := u.Authenticate(test.args.ctx, test.args.input)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
