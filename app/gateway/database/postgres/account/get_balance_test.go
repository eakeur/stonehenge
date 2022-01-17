package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/gateway/database/postgres/postgrestest"
	"testing"
)

func TestGetBalance(t *testing.T) {

	db, err := postgrestest.NewCleanDatabase()
	if err != nil {
		t.Fatalf("could not get database: %v", err)
	}

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
				accounts, err := postgrestest.PopulateAccounts(test.args.ctx, postgrestest.GetFakeAccount())
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
			defer postgrestest.RecycleDatabase(test.args.ctx)
			repo := NewRepository(db)

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
