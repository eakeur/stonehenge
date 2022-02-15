package transfer

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/gateway/postgres/tests"
	"testing"
)

func TestList(t *testing.T) {
	t.Parallel()
	log := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

	type args struct {
		ctx context.Context
	}

	type test struct {
		name    string
		args    args
		before  func(test test, db tests.Database) (transfer.Filter, error)
		want    []transfer.Transfer
		wantErr error
	}

	cases := []test{

		// Should return the account expected
		{
			name: "return account expected",
			before: func(test test, db tests.Database) (transfer.Filter, error) {
				accounts, err := db.PopulateAccounts(test.args.ctx, account.GetFakeAccounts()...)
				if err != nil {
					return transfer.Filter{}, err
				}

				_, err = db.PopulateTransfers(test.args.ctx, transfer.GetFakeTransfers()...)

				if err != nil {
					return transfer.Filter{}, err
				}

				return transfer.Filter{
					OriginID: accounts[0].ExternalID,
				}, nil
			},
			args: args{ctx: context.Background()},
			want: []transfer.Transfer{
				{
					OriginID:      1,
					DestinationID: 2,
					Amount:        500,
				},
				{
					OriginID:      1,
					DestinationID: 2,
					Amount:        500,
				},
				{
					OriginID:      1,
					DestinationID: 2,
					Amount:        500,
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

			filter, err := test.before(test, db)
			if err != nil {
				t.Fatalf("error running routine before: %v", err)
			}

			got, err := repo.List(test.args.ctx, filter)

			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, len(test.want), len(got))
			if len(test.want) == len(got) {
				for i, tr := range got {
					exp := test.want[i]
					assert.Equal(t, filter.OriginID, tr.Details.OriginExternalID)
					assert.Equal(t, exp.OriginID, tr.OriginID)
					assert.Equal(t, exp.DestinationID, tr.DestinationID)
					assert.Equal(t, exp.Amount, tr.Amount)
					assert.NotNil(t, tr.CreatedAt)
					assert.NotNil(t, tr.UpdatedAt)
					assert.NotNil(t, tr.ID)
					assert.NotNil(t, tr.ExternalID)
				}
			}
		})
	}

}
