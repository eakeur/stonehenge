package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/password"
	"stonehenge/app/gateway/database/postgres/postgrestest"
	"testing"
)


func TestGetWithExternalID(t *testing.T) {

	db, err := postgrestest.NewCleanDatabase()
	if err != nil {
		t.Fatalf("could not get database: %v", err)
	}

	type args struct {
		ctx      context.Context
		document document.Document
	}

	type test struct {
		name    string
		args    args
		before  func(test, account.Repository) (account.Account, error)
		want    account.Account
		wantErr error
	}

	tests := []test{

		// Should return the account expected
		{
			name: "return account expected",
			before: func(test test, a account.Repository) (account.Account, error) {
				accounts, err := postgrestest.PopulateAccounts(test.args.ctx, account.Account{
					Document: "05161964057",
					Secret:   password.From("12345678"),
					Name:     "Spencer Reis",
					Balance:  758250,
				})
				return accounts[0], err
			},
			args: args{ctx: context.Background(), document: "05161964057"},
			want: account.Account{},
		},

		// Should return ErrNotFound for unknown document
		{
			name:    "return ErrNotFound unknown document",
			args:    args{ctx: context.Background(), document: "05161964057"},
			want:    account.Account{},
			wantErr: account.ErrNotFound,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			defer postgrestest.RecycleDatabase(test.args.ctx)
			repo := NewRepository(db)

			if test.before != nil {
				acc, err := test.before(test, repo)
				test.want = acc
				if err != nil {
					t.Fatalf("error running routine before: %v", err)
				}
			}

			got, err := repo.GetByExternalID(test.args.ctx, test.want.ExternalID)

			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
