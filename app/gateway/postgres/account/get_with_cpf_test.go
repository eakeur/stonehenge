package account

import (
	"context"
	"os"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/document"
	"stonehenge/app/gateway/postgres/tests"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestGetWithCPF(t *testing.T) {

	t.Parallel()
	log := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

	type args struct {
		ctx      context.Context
		document document.Document
	}

	type test struct {
		name    string
		args    args
		before  func(test, tests.Database) (account.Account, error)
		want    account.Account
		wantErr error
	}

	cases := []test{

		// Should return the account expected
		{
			name: "return account expected",
			before: func(test test, db tests.Database) (account.Account, error) {
				accounts, err := db.PopulateAccounts(test.args.ctx, account.GetFakeAccount())
				return accounts[0], err
			},
			args: args{ctx: context.Background(), document: "24788516002"},
			want: account.GetFakeAccount(),
		},

		// Should return ErrNotFound for unknown document
		{
			name:    "return ErrNotFound unknown document",
			args:    args{ctx: context.Background(), document: "05161964057"},
			want:    account.Account{},
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

			if test.before != nil {
				acc, err := test.before(test, db)
				if err != nil {
					t.Fatalf("error running routine before: %v", err)
				}
				test.want.ID = acc.ID
				test.want.ExternalID = acc.ExternalID
				test.want.CreatedAt = acc.CreatedAt
				test.want.UpdatedAt = acc.UpdatedAt
			}

			got, err := repo.GetWithCPF(test.args.ctx, test.args.document)

			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
