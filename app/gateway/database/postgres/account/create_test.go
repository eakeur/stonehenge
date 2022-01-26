package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/config"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/gateway/database/postgres/postgrestest"
	"stonehenge/app/gateway/database/postgres/transaction"
	"stonehenge/app/gateway/logger"
	"testing"
)

func TestCreate(t *testing.T) {
	db, err := postgrestest.NewCleanDatabase()
	if err != nil {
		t.Fatalf("could not get database: %v", err)
	}

	tx := transaction.NewManager(db)

	log := logger.NewLogger(config.LoggerConfigurations{Environment: "development"})

	type args struct {
		ctx     context.Context
		account account.Account
	}

	type test struct {
		name    string
		args    args
		before  func(test) error
		want    account.Account
		wantErr error
	}

	tests := []test{

		// Should return Account with generated fields
		{
			name: "return Account with generated fields",
			args: args{ctx: context.Background(), account: postgrestest.GetFakeAccount()},
			want: postgrestest.GetFakeAccount(),
		},

		// Should return ErrAlreadyExist for duplicate account
		{
			name: "return ErrAlreadyExist for duplicate account",
			before: func(test test) error {
				_, err := postgrestest.PopulateAccounts(test.args.ctx, postgrestest.GetFakeAccount())
				return err
			},
			args:    args{ctx: context.Background(), account: postgrestest.GetFakeAccount()},
			wantErr: account.ErrAlreadyExist,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			defer postgrestest.RecycleDatabase(test.args.ctx)

			repo := NewRepository(db, log)

			if test.before != nil {
				err := test.before(test)
				if err != nil {
					t.Fatalf("error running routine before: %v", err)
				}
			}
			ctx, err := tx.Begin(test.args.ctx)
			if err != nil {
				t.Fatalf("could not start transaction: %v", err)
			}
			defer tx.Rollback(ctx)

			acc, err := repo.Create(ctx, test.args.account)
			if err == nil {
				if err := tx.Commit(ctx); err != nil {
					t.Fatalf("could not commit transaction: %v", err)
				}

				test.want.ID = acc.ID
				test.want.ExternalID = acc.ExternalID
				test.want.CreatedAt = acc.CreatedAt
				test.want.UpdatedAt = acc.UpdatedAt
			}

			assert.ErrorIs(t, test.wantErr, err)
			assert.Equal(t, test.want, acc)
		})
	}
}
