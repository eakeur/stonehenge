package account

import (
	"context"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountCreation(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	accountID := id.ExternalFrom("d0052623-0695-4a3a-abf6-887f613dda8e")

	singleInstancePassword := password.From("D@V@C@O@")

	transactions := &transaction.RepositoryMock{BeginResult: ctx}

	accesses := &access.RepositoryMock{
		CreateResult: access.Access{AccountID: accountID},
	}

	type args struct {
		ctx   context.Context
		input CreateInput
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
		want    CreateOutput
		wantErr error
	}

	tests := []test{

		{
			name: "return ErrTokenFailedCreation for unsuccessful token encryption",
			fields: fields{repo: &account.RepositoryMock{
				CreateResult: account.Account{
					ExternalID: accountID,
					Document:   "97662062015",
					Secret:     singleInstancePassword,
					Name:       "Lina Pereira",
					Balance:    5000,
				}},
				access: access.RepositoryMock{
					Error: access.ErrTokenFailedCreation,
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				Document: "97662062015",
				Secret:   singleInstancePassword,
				Name:     "Lina Pereira",
			}},
			wantErr: access.ErrTokenFailedCreation,
		},

		{
			name: "return CreateOutput for created account",
			fields: fields{repo: &account.RepositoryMock{
				CreateResult: account.Account{
					ExternalID: accountID,
					Document:   "97662062015",
					Secret:     singleInstancePassword,
					Name:       "Lina Pereira",
					Balance:    5000,
				}},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				Document: "97662062015",
				Secret:   singleInstancePassword,
				Name:     "Lina Pereira",
			}},
			want: CreateOutput{
				AccountID: accountID,
				Access:    access.Access{AccountID: accountID}},
		},

		{
			name:   "return ErrAlreadyExists for duplicate document",
			fields: fields{repo: &account.RepositoryMock{Error: account.ErrAlreadyExist}},
			args: args{ctx: context.Background(), input: CreateInput{
				Document: "97662062015",
				Secret:   singleInstancePassword,
				Name:     "Lina Pereira",
			}},
			wantErr: account.ErrAlreadyExist,
		},

		{
			name:   "return ErrInvalidDocument for corrupted CPF",
			fields: fields{repo: &account.RepositoryMock{}},
			args: args{ctx: context.Background(), input: CreateInput{
				Document: "9766206201",
				Secret:   singleInstancePassword,
				Name:     "Lina Pereira",
			}},
			wantErr: document.ErrInvalidDocument,
		},

		{
			name:   "return ErrInvalidDocument for valid CPF with symbols",
			fields: fields{repo: &account.RepositoryMock{}},
			args: args{ctx: context.Background(), input: CreateInput{
				Document: "976.620.620-10",
				Secret:   password.From("D@V@C@O@"),
				Name:     "Lina Pereira",
			}},
			wantErr: document.ErrInvalidDocument,
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
			got, err := u.Create(test.args.ctx, test.args.input)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
