package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/config"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/document"
	"stonehenge/app/gateway/database/postgres/postgrestest"
	"stonehenge/app/gateway/logger"
	"testing"
)

func TestGetWithCPF(t *testing.T) {

	db, err := postgrestest.NewCleanDatabase()
	if err != nil {
		t.Fatalf("could not get database: %v", err)
	}
	log := logger.NewLogger(config.LoggerConfigurations{Environment: "development"})

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
				accounts, err := postgrestest.PopulateAccounts(test.args.ctx, postgrestest.GetFakeAccount())
				return accounts[0], err
			},
			args: args{ctx: context.Background(), document: "24788516002"},
			want: postgrestest.GetFakeAccount(),
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
			repo := NewRepository(db, log)

			if test.before != nil {
				acc, err := test.before(test)
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
