package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/password"
	"stonehenge/app/gateway/database/postgres/postgrestest"
	"testing"
)

func TestList(t *testing.T) {

	db, err := postgrestest.NewCleanDatabase()
	if err != nil {
		t.Fatalf("could not get database: %v", err)
	}

	type args struct {
		ctx    context.Context
		filter account.Filter
	}

	type test struct {
		name    string
		args    args
		before  func(test, account.Repository) ([]account.Account, error)
		want    []account.Account
		wantErr error
	}

	tests := []test{

		// Should return the account expected
		{
			name: "return account expected",
			before: func(test test, a account.Repository) ([]account.Account, error) {
				accounts, err := postgrestest.PopulateAccounts(test.args.ctx,
					account.Account{
						Document: "70830052062",
						Secret:   password.From("12345678"),
						Name:     "John Reis",
						Balance:  2500,
					},
					account.Account{
						Document: "24388516007",
						Secret:   password.From("12345678"),
						Name:     "Wagner Reis",
						Balance:  4500,
					},
					account.Account{
						Document: "05161964057",
						Secret:   password.From("12345678"),
						Name:     "Spencer Reis",
						Balance:  5000,
					},
					account.Account{
						Document: "24788516002",
						Secret:   password.From("12345678"),
						Name:     "Lina Pereira",
						Balance:  4500,
					},
					account.Account{
						Document: "24385516005",
						Secret:   password.From("12345678"),
						Name:     "Elza Soares",
						Balance:  4500,
					},
					account.Account{
						Document: "24384516008",
						Secret:   password.From("12345678"),
						Name:     "Jur Arras",
						Balance:  4500,
					})
				return accounts, err
			},
			args: args{ctx: context.Background(), filter: account.Filter{Name: "Reis"}},
			want: []account.Account{
				{
					Document: "70830052062",
					Secret:   password.From("12345678"),
					Name:     "John Reis",
					Balance:  2500,
				},
				{
					Document: "24388516007",
					Secret:   password.From("12345678"),
					Name:     "Wagner Reis",
					Balance:  4500,
				},
				{
					Document: "05161964057",
					Secret:   password.From("12345678"),
					Name:     "Spencer Reis",
					Balance:  5000,
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			defer postgrestest.RecycleDatabase(test.args.ctx)
			repo := NewRepository(db)

			if test.before != nil {
				_, err := test.before(test, repo)
				if err != nil {
					t.Fatalf("error running routine before: %v", err)
				}
			}

			got, err := repo.List(test.args.ctx, test.args.filter)

			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, len(test.want), len(got))
			if len(test.want) == len(got) {
				for i, acc := range got {
					exp := test.want[i]
					assert.Equal(t, exp.Name, acc.Name)
					assert.Equal(t, exp.Document, acc.Document)
					assert.Equal(t, exp.Balance, acc.Balance)
				}
			}
		})
	}

}
