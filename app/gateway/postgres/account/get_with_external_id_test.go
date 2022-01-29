package account

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/document"
	postgrestest2 "stonehenge/app/gateway/postgres/postgrestest"
	"testing"
)

func TestGetWithExternalID(t *testing.T) {

	db, err := postgrestest2.NewCleanDatabase()
	if err != nil {
		t.Fatalf("could not get database: %v", err)
	}

	log := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

	type args struct {
		ctx      context.Context
		document document.Document
	}

	type test struct {
		name    string
		args    args
		before  func(test) (account.Account, error)
		want    account.Account
		wantErr error
	}

	tests := []test{

		// Should return the account expected
		{
			name: "return account expected",
			before: func(test test) (account.Account, error) {
				accounts, err := postgrestest2.PopulateAccounts(test.args.ctx, postgrestest2.GetFakeAccount())
				return accounts[0], err
			},
			args: args{ctx: context.Background(), document: "05161964057"},
			want: postgrestest2.GetFakeAccount(),
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
			defer postgrestest2.RecycleDatabase(test.args.ctx)
			repo := NewRepository(db, log)

			if test.before != nil {
				acc, err := test.before(test)
				test.want = acc
				if err != nil {
					t.Fatalf("error running routine before: %v", err)
				}
				test.want.ID = acc.ID
				test.want.ExternalID = acc.ExternalID
				test.want.CreatedAt = acc.CreatedAt
				test.want.UpdatedAt = acc.UpdatedAt
			}

			got, err := repo.GetByExternalID(test.args.ctx, test.want.ExternalID)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
