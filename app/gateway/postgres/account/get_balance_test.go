package account

import (
	"context"
	"os"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/gateway/postgres/tests"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {

	t.Parallel()

	log := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	type args struct {
		ctx context.Context
	}

	type test struct {
		name    string
		args    args
		before  func(test, tests.Database) (account.Account, error)
		want    currency.Currency
		wantErr error
	}

	cases := []test{

		// Should return the balance expected for this account
		{
			name: "return balance expected",
			before: func(test test, db tests.Database) (account.Account, error) {
				accounts, err := db.PopulateAccounts(test.args.ctx, account.GetFakeAccount())
				return accounts[0], err
			},
			args: args{ctx: context.Background()},
			want: 50000000,
		},

		// Should return ErrNotFound for nonexistent account on database
		{
			name: "return ErrNotFound for nonexistent account",
			before: func(test test, db tests.Database) (account.Account, error) {
				return account.Account{}, nil
			},
			args:    args{ctx: context.Background()},
			want:    0,
			wantErr: account.ErrNotFound,
		},
	}

	for _, test := range cases {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db := tests.NewTestDatabase(t)
			defer db.Drop()

			repo := NewRepository(db.Pool, log)

			acc, err := test.before(test, db)
			if err != nil {
				t.Fatalf("error running routine before: %v", err)
			}

			got, err := repo.GetBalance(test.args.ctx, acc.ExternalID)

			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
