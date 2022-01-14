package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"stonehenge/app/gateway/database/postgres/postgrestest"
	"testing"
)

func TestGetBalance(t *testing.T) {

	db, err := postgrestest.NewCleanDatabase()
	if err != nil {
		t.Fatalf("could not get database: %v", err)
	}

	type args struct {
		ctx     context.Context
	}

	type test struct {
		name    string
		args    args
		before  func(test, account.Repository) (id.ExternalID, error)
		want	currency.Currency
		wantErr error
	}

	tests := []test{

		// Should return the balance expected for this account
		{
			name: "return balance expected",
			before: func(test test, a account.Repository) (id.ExternalID, error) {
				accounts, err := postgrestest.PopulateAccounts(test.args.ctx, account.Account{
					Document: "05161964057",
					Secret:   password.From("12345678"),
					Name:     "Spencer Reis",
					Balance:  758250,
				})
				return accounts[0].ExternalID, err
			},
			args: args{ctx: context.Background()},
			want: 758250,
		},

		// Should return ErrNotFound for nonexistent account on database
		{
			name: "return ErrNotFound for nonexistent account",
			before: func(test test, a account.Repository) (id.ExternalID, error) {
				return id.New(), nil
			},
			args: args{ctx: context.Background()},
			want: 0,
			wantErr: account.ErrNotFound,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			defer postgrestest.RecycleDatabase(test.args.ctx)
			repo := NewRepository(db)

			ext, err := test.before(test, repo)
			if err != nil {
				t.Fatalf("error running routine before: %v", err)
			}

			got, err := repo.GetBalance(test.args.ctx, ext)

			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}


}
