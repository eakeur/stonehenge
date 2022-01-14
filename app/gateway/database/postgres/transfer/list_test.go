package transfer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
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
		ctx context.Context
	}

	type test struct {
		name    string
		args    args
		before  func(test) (transfer.Filter, error)
		want    []transfer.Transfer
		wantErr error
	}

	tests := []test{

		// Should return the account expected
		{
			name: "return account expected",
			before: func(test test) (transfer.Filter, error) {
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
					})
				if err != nil {
					return transfer.Filter{}, err
				}

				_, err = postgrestest.PopulateTransfers(test.args.ctx,
					transfer.Transfer{
						OriginID:      1,
						DestinationID: 2,
						Amount:        500,
					},
					transfer.Transfer{
						OriginID:      1,
						DestinationID: 2,
						Amount:        500,
					},
					transfer.Transfer{
						OriginID:      2,
						DestinationID: 1,
						Amount:        500,
					},
					transfer.Transfer{
						OriginID:      2,
						DestinationID: 1,
						Amount:        500,
					},
					transfer.Transfer{
						OriginID:      1,
						DestinationID: 2,
						Amount:        500,
					})

				if err != nil {
					return transfer.Filter{}, err
				}

				return transfer.Filter{
					OriginID: accounts[0].ExternalID,
				}, err
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

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			defer postgrestest.RecycleDatabase(test.args.ctx)
			repo := NewRepository(db)

			filter, err := test.before(test)
			if err != nil {
				t.Fatalf("error running routine before: %v", err)
			}

			got, err := repo.List(test.args.ctx, filter)

			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, len(test.want), len(got))
			if len(test.want) == len(got) {
				for _, acc := range got {
					assert.Equal(t, filter.OriginID, acc.Details.OriginExternalID)
				}
			}
		})
	}

}
