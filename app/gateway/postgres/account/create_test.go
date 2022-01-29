package account

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"stonehenge/app/core/entities/account"
	postgrestest2 "stonehenge/app/gateway/postgres/postgrestest"
	"stonehenge/app/gateway/postgres/transaction"
	"testing"
)

func TestCreate(t *testing.T) {
	db, err := postgrestest2.NewCleanDatabase()
	if err != nil {
		t.Fatalf("could not get database: %v", err)
	}

	tx := transaction.NewManager(db)

	log := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

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
			args: args{ctx: context.Background(), account: postgrestest2.GetFakeAccount()},
			want: postgrestest2.GetFakeAccount(),
		},

		// Should return ErrAlreadyExist for duplicate account
		{
			name: "return ErrAlreadyExist for duplicate account",
			before: func(test test) error {
				_, err := postgrestest2.PopulateAccounts(test.args.ctx, postgrestest2.GetFakeAccount())
				return err
			},
			args:    args{ctx: context.Background(), account: postgrestest2.GetFakeAccount()},
			wantErr: account.ErrAlreadyExist,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			defer postgrestest2.RecycleDatabase(test.args.ctx)

			repo := NewRepository(db, log)

			if test.before != nil {
				err := test.before(test)
				if err != nil {
					t.Fatalf("error running routine before: %v", err)
				}
			}
			ctx := tx.Begin(test.args.ctx)
			defer tx.End(ctx)

			acc, err := repo.Create(ctx, test.args.account)
			if err == nil {
				test.want.ID = acc.ID
				test.want.ExternalID = acc.ExternalID
				test.want.CreatedAt = acc.CreatedAt
				test.want.UpdatedAt = acc.UpdatedAt
			}

			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, acc)
		})
	}
}
