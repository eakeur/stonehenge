package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/password"
	"stonehenge/app/gateway/database/postgres/postgres_test"
	"stonehenge/app/gateway/database/postgres/transaction"
	"testing"
)

func TestCreate(t *testing.T) {

	t.Parallel()

	db := postgres_test.GetDB()

	tx := transaction.NewTransaction(db)


	type args struct {
		ctx     context.Context
		account account.Account
	}

	type test struct {
		name    string
		args    args
		wantErr error
	}

	tests := []test{

		// Should return Account with generated fields
		{
			name: "return Account with generated fields",
			args: args{ctx: context.Background(), account: account.Account{
				Document: "05161964057",
				Secret:   password.From("12345678"),
				Name:     "Spencer Reis",
				Balance:  5000,
			}},

		},

		// Should return ErrAlreadyExist for duplicate account
		{
			name: "return ErrAlreadyExist for duplicate account",
			args: args{ctx: context.Background(), account: account.Account{
				Document: "05161964057",
				Secret:   password.From("12345678"),
				Name:     "Spencer Reis",
				Balance:  5000,
			}},
			wantErr: account.ErrAlreadyExist,

		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			repo := NewRepository(db)
			ctx, err := tx.Begin(test.args.ctx)
			if err != nil {
				t.Error(err)
			}

			acc, err := repo.Create(ctx, test.args.account)
			if err != nil {
				tx.Rollback(ctx)
				assert.ErrorIs(t, test.wantErr, err)
				return
			}

			if err := tx.Commit(ctx); err != nil {
				tx.Rollback(ctx)
				assert.ErrorIs(t, test.wantErr, err)
			}

			accountInDB, err := repo.GetByExternalID(ctx, acc.ExternalID)

			assert.ErrorIs(t, test.wantErr, err)
			assert.Equal(t, accountInDB, acc)
		})
	}
}
