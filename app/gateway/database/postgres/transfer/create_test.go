package transfer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"stonehenge/app/gateway/database/postgres/postgrestest"
	"stonehenge/app/gateway/database/postgres/transaction"
	"testing"
)

func TestCreate(t *testing.T) {
	db, err := postgrestest.NewCleanDatabase()
	if err != nil {
		t.Fatalf("could not get database: %v", err)
	}

	tx := transaction.NewTransaction(db)

	type args struct {
		ctx      context.Context
		transfer transfer.Transfer
	}

	type test struct {
		name    string
		args    args
		before  func(test) error
		wantErr error
	}

	tests := []test{

		// Should return Transfer with generated fields
		{
			name: "return Transfer with generated fields",
			before: func(test test) error {
				_, err := postgrestest.PopulateAccounts(test.args.ctx,
					account.Account{
						Document: "05161964057",
						Secret:   password.From("12345678"),
						Name:     "Spencer Reis",
						Balance:  5000,
					},
					account.Account{
						Document: "05161964056",
						Secret:   password.From("87654321"),
						Name:     "Nina Simone",
						Balance:  50000,
					})
				if err != nil {
					return err
				}
				return nil
			},
			args: args{ctx: context.Background(), transfer: transfer.Transfer{
				OriginID:      1,
				DestinationID: 2,
				Amount:        500,
			}},
		},

		// Should return ErrNonexistentOrigin for unknown origin id
		{
			name: "return ErrNonexistentOrigin ",
			before: func(test test) error {
				_, err := postgrestest.PopulateAccounts(test.args.ctx,
					account.Account{
						Document: "05161964056",
						Secret:   password.From("87654321"),
						Name:     "Nina Simone",
						Balance:  50000,
					})
				if err != nil {
					return err
				}
				return nil
			},
			args: args{ctx: context.Background(), transfer: transfer.Transfer{
				OriginID:      2,
				DestinationID: 1,
				Amount:        500,
			}},
			wantErr: transfer.ErrNonexistentOrigin,
		},

		// Should return ErrNonexistentOrigin for unknown origin id
		{
			name: "return ErrNonexistentDestination",
			before: func(test test) error {
				_, err := postgrestest.PopulateAccounts(test.args.ctx,
					account.Account{
						Document: "05161964056",
						Secret:   password.From("87654321"),
						Name:     "Nina Simone",
						Balance:  50000,
					})
				if err != nil {
					return err
				}
				return nil
			},
			args: args{ctx: context.Background(), transfer: transfer.Transfer{
				OriginID:      1,
				DestinationID: 3,
				Amount:        500,
			}},
			wantErr: transfer.ErrNonexistentDestination,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			defer postgrestest.RecycleDatabase(test.args.ctx)

			repo := NewRepository(db)

			if test.before != nil {
				err := test.before(test)
				if err != nil {
					t.Fatalf("error running routine before: %v", err)
				}
			}
			ctx, err := tx.Begin(test.args.ctx)
			if err != nil {
				t.Error(err)
			}

			tr, err := repo.Create(ctx, test.args.transfer)
			if err != nil {
				tx.Rollback(ctx)
				assert.ErrorIs(t, test.wantErr, err)
				return
			}

			if err := tx.Commit(ctx); err != nil {
				tx.Rollback(ctx)
				assert.ErrorIs(t, test.wantErr, err)
			}

			assert.ErrorIs(t, test.wantErr, err)
			if test.wantErr != nil {
				assert.Equal(t, id.ID(1), tr.ID)
			}
		})
	}
}
