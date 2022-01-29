package account

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	postgrestest2 "stonehenge/app/gateway/postgres/postgrestest"
	"testing"
)

func TestGetBalance(t *testing.T) {

	db, err := postgrestest2.NewCleanDatabase()
	if err != nil {
		t.Fatalf("could not get database: %v", err)
	}

	log := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	type args struct {
		ctx context.Context
	}

	type test struct {
		name    string
		args    args
		before  func(test) (account.Account, error)
		want    currency.Currency
		wantErr error
	}

	tests := []test{

		// Should return the balance expected for this account
		{
			name: "return balance expected",
			before: func(test test) (account.Account, error) {
				accounts, err := postgrestest2.PopulateAccounts(test.args.ctx, postgrestest2.GetFakeAccount())
				return accounts[0], err
			},
			args: args{ctx: context.Background()},
			want: 4500,
		},

		// Should return ErrNotFound for nonexistent account on database
		{
			name: "return ErrNotFound for nonexistent account",
			before: func(test test) (account.Account, error) {
				return account.Account{}, nil
			},
			args:    args{ctx: context.Background()},
			want:    0,
			wantErr: account.ErrNotFound,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			defer postgrestest2.RecycleDatabase(test.args.ctx)
			repo := NewRepository(db, log)

			acc, err := test.before(test)
			if err != nil {
				t.Fatalf("error running routine before: %v", err)
			}

			got, err := repo.GetBalance(test.args.ctx, acc.ExternalID)

			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
