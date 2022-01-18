package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"testing"
)

func TestAccountCreation(t *testing.T) {
	t.Parallel()

	tx := &transaction.RepositoryMock{
		BeginFunc: func(ctx context.Context) (context.Context, error) {
			return ctx, nil
		},
	}

	tk := &access.RepositoryMock{
		CreateResult: access.Access{
			AccountID: id.ExternalFrom(accountID),
		},
	}

	type args struct {
		ctx   context.Context
		input CreateInput
	}

	type fields struct {
		tx   transaction.Transaction
		tk   access.Manager
		repo account.Repository
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    CreateOutput
		wantErr error
	}

	tests := []test{

		// Should return CreateOutput with filled fields for successfully created account
		{
			name: "return CreateOutput for created account",
			fields: fields{tx: tx, tk: tk, repo: &account.RepositoryMock{
				CreateFunc: func(ctx context.Context, account account.Account) (account.Account, error) {
					account.ExternalID = id.ExternalFrom(accountID)
					return account, nil
				},
			}},
			args: args{ctx: context.Background(), input: CreateInput{
				Document: "97662062015",
				Secret:   password.From("D@V@C@O@"),
				Name:     "Lina Pereira",
			}},
			want: CreateOutput{AccountID: id.ExternalFrom(accountID), Access: access.Access{AccountID: id.ExternalFrom(accountID)}},
		},

		// Should return ErrAlreadyExists for document related to another account already
		{
			name:   "return ErrAlreadyExists for duplicate document",
			fields: fields{tx: tx, tk: tk, repo: &account.RepositoryMock{Error: account.ErrAlreadyExist}},
			args: args{ctx: context.Background(), input: CreateInput{
				Document: "97662062015",
				Secret:   password.From("D@V@C@O@"),
				Name:     "Lina Pereira",
			}},
			wantErr: account.ErrAlreadyExist,
		},

		// Should return ErrInvalidDocument for applying with corrupted CPF
		{
			name:   "return ErrInvalidDocument for corrupted CPF",
			fields: fields{tx: tx, tk: tk, repo: &account.RepositoryMock{}},
			args: args{ctx: context.Background(), input: CreateInput{
				Document: "9766206201",
				Secret:   password.From("D@V@C@O@"),
				Name:     "Lina Pereira",
			}},
			wantErr: document.ErrInvalidDocument,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			u := New(test.fields.repo, test.fields.tx, test.fields.tk)
			got, err := u.Create(test.args.ctx, test.args.input)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
