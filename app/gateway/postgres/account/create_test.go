package account

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/gateway/postgres/tests"
	"stonehenge/app/gateway/postgres/transaction"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	log := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

	type args struct {
		ctx     context.Context
		account account.Account
	}

	type test struct {
		name    string
		args    args
		before  func(test, tests.Database) error
		want    account.Account
		wantErr error
	}

	cases := []test{

		// Should return Account with generated fields
		{
			name: "return Account with generated fields",
			args: args{ctx: context.Background(), account: account.GetFakeAccount()},
			want: account.GetFakeAccount(),
		},

		// Should return ErrAlreadyExist for duplicate account
		{
			name: "return ErrAlreadyExist for duplicate account",
			before: func(test test, db tests.Database) error {
				_, err := db.PopulateAccounts(test.args.ctx, account.GetFakeAccount())
				return err
			},
			args:    args{ctx: context.Background(), account: account.GetFakeAccount()},
			wantErr: account.ErrAlreadyExist,
		},
	}

	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db := tests.NewTestDatabase(t)
			defer db.Drop()

			repo := NewRepository(db.Pool, log)
			tx := transaction.NewManager(db.Pool)

			if test.before != nil {
				err := test.before(test, db)
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
