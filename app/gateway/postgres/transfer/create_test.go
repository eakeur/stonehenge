package transfer

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/gateway/postgres/tests"
	"stonehenge/app/gateway/postgres/transaction"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	log := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

	type args struct {
		ctx      context.Context
		transfer transfer.Transfer
	}

	type test struct {
		name    string
		args    args
		before  func(test test, db tests.Database) error
		want    transfer.Transfer
		wantErr error
	}

	cases := []test{

		// Should return Transfer with generated fields
		{
			name: "return Transfer with generated fields",
			before: func(test test, db tests.Database) error {
				_, err := db.PopulateAccounts(test.args.ctx, account.GetFakeAccounts()...)
				return err
			},
			args: args{ctx: context.Background(), transfer: transfer.Transfer{
				OriginID:      1,
				DestinationID: 2,
				Amount:        500,
				EffectiveDate: func() time.Time {
					t, _ := time.Parse("2006/01/02 15:04:05", "2022/01/01 10:50:01")
					return t
				}(),
			}},
			want: transfer.Transfer{
				OriginID:      1,
				DestinationID: 2,
				Amount:        500,
				EffectiveDate: func() time.Time {
					t, _ := time.Parse("2006/01/02 15:04:05", "2022/01/01 10:50:01")
					return t
				}(),
			},
		},

		// Should return ErrNonexistentOrigin for unknown origin id
		{
			name: "return ErrNonexistentOrigin ",
			before: func(test test, db tests.Database) error {
				_, err := db.PopulateAccounts(test.args.ctx, account.GetFakeAccount())
				return err
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
			before: func(test test, db tests.Database) error {
				_, err := db.PopulateAccounts(test.args.ctx, account.GetFakeAccount())
				return err
			},
			args: args{ctx: context.Background(), transfer: transfer.Transfer{
				OriginID:      1,
				DestinationID: 3,
				Amount:        500,
			}},
			wantErr: transfer.ErrNonexistentDestination,
		},
	}

	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db := tests.NewTestDatabase(t)
			defer db.Drop()

			repo := NewRepository(db.Pool, log)
			tx := transaction.NewManager(db.Pool)

			if test.before != nil {
				err := test.before(test, db)
				if err != nil {
					t.Fatalf("error running routine before: %v", err)
				}
			}
			ctx := tx.Begin(test.args.ctx)
			defer tx.End(ctx)

			tr, err := repo.Create(ctx, test.args.transfer)
			if err == nil {
				test.want.ID = tr.ID
				test.want.ExternalID = tr.ExternalID
				test.want.CreatedAt = tr.CreatedAt
				test.want.UpdatedAt = tr.UpdatedAt
			}

			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, tr)
		})
	}
}
