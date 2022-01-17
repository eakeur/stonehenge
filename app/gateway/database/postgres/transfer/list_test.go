package transfer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/transfer"
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
				accounts, err := postgrestest.PopulateAccounts(test.args.ctx, postgrestest.GetFakeAccounts()...)
				if err != nil {
					return transfer.Filter{}, err
				}

				_, err = postgrestest.PopulateTransfers(test.args.ctx, postgrestest.GetFakeTransfers()...)

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
