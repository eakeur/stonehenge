package account

import (
	"context"
	"os"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/gateway/postgres/tests"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {

	t.Parallel()

	log := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

	type args struct {
		ctx    context.Context
		filter account.Filter
	}

	type test struct {
		name    string
		args    args
		before  func(test, tests.Database) ([]account.Account, error)
		want    []account.Account
		wantErr error
	}

	cases := []test{

		// Should return the account expected
		{
			name: "return account expected",
			before: func(test test, db tests.Database) ([]account.Account, error) {
				return db.PopulateAccounts(test.args.ctx, account.GetFakeAccounts()...)
			},
			args: args{ctx: context.Background(), filter: account.Filter{Name: "Reis"}},
			want: []account.Account{
				{
					Document: "70830052062",
					Name:     "John Reis",
					Balance:  105000,
				},
				{
					Document: "24388516007",
					Name:     "Wagner Reis",
					Balance:  450000,
				},
				{
					Document: "05161964057",
					Name:     "Spencer Reis",
					Balance:  502200,
				},
			},
		},
	}

	for _, test := range cases {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db := tests.NewTestDatabase(t)
			defer db.Drop()

			repo := NewRepository(db.Pool, log)

			if test.before != nil {
				_, err := test.before(test, db)
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
					assert.Nil(t, acc.Secret.Compare("12345678"))
					assert.NotNil(t, acc.CreatedAt)
					assert.NotNil(t, acc.UpdatedAt)
					assert.NotNil(t, acc.ID)
					assert.NotNil(t, acc.ExternalID)
				}
			}
		})
	}

}
