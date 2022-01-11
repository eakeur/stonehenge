package transfer

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"testing"

	"github.com/golang-migrate/migrate/v4/database/pgx"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx   context.Context
		input CreateInput
	}

	type fields struct {
		ac account.Repository
		tr transfer.Repository
		tx transaction.Transaction
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    CreateOutput
		wantErr error
	}

	const (
		idOrigin      string = "d0052623-0695-4a3a-abf6-887f613dda8e"
		idDestination string = "17edb329-4b65-41ba-bb26-5060a1e157ab"
		idTransfer    string = "b8d11928-3eab-45a8-8be3-31411bd120f2"
	)

	tests := []test{

		// successful creation of transfer

		{
			name: "should return successful creation of transfer",
			fields: fields{
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.ExternalID) (account.Account, error) {
						internal := 1
						if ext == id.From(idDestination) {
							internal = 2
						}
						return account.Account{
							ID:      id.ID(internal),
							Balance: 5000,
						}, nil
					},
					UpdateBalanceFunc: func(ctx context.Context, id id.ExternalID, balance currency.Currency) error {
						return nil
					},
				},
				tr: &transfer.RepositoryMock{
					CreateFunc: func(ctx context.Context, transfer transfer.Transfer) (transfer.Transfer, error) {
						transfer.ExternalID = id.From(idTransfer)
						return transfer, nil
					},
				},
				tx: &transaction.RepositoryMock{
					BeginFunc: func(ctx context.Context) (context.Context, error) {
						return ctx, nil
					},
					CommitFunc: func(ctx context.Context) error {
						return nil
					},
					RollbackFunc: func(ctx context.Context) {},
				},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					OriginID: id.From(idOrigin),
					DestID:   id.From(idDestination),
					Amount:   2500,
				},
			},
			want: CreateOutput{
				RemainingBalance: 2500,
				TransferId:       id.From(idTransfer),
			},
		},

		// not enough money in origin balance
		{
			name: "should return unsuccessful creation of transfer due to not enough money in origin balance",
			fields: fields{
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.ExternalID) (account.Account, error) {
						internal := 1
						if ext == id.From(idDestination) {
							internal = 2
						}
						return account.Account{
							ID:      id.ID(internal),
							Balance: 1250,
						}, nil
					},
					UpdateBalanceFunc: func(ctx context.Context, id id.ExternalID, balance currency.Currency) error {
						return nil
					},
				},
				tr: &transfer.RepositoryMock{},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					OriginID: id.From(idOrigin),
					DestID:   id.From(idDestination),
					Amount:   2500,
				},
			},
			want:    CreateOutput{},
			wantErr: account.ErrNoMoney,
		},

		// operation being made to the same account of origin
		{
			name: "should return unsuccessful creation of transfer due to operation being made to the same account of origin",
			fields: fields{
				ac: &account.RepositoryMock{},
				tr: &transfer.RepositoryMock{},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					OriginID: id.From(idOrigin),
					DestID:   id.From(idOrigin),
					Amount:   2500,
				},
			},
			want:    CreateOutput{},
			wantErr: transfer.ErrSameAccount,
		},

		// amount in transfer is not valid
		{
			name: "should return unsuccessful creation of transfer due to invalid amount on transfer request",
			fields: fields{
				ac: &account.RepositoryMock{},
				tr: &transfer.RepositoryMock{},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					OriginID: id.From(idOrigin),
					DestID:   id.From(idDestination),
					Amount:   0,
				},
			},
			want:    CreateOutput{},
			wantErr: transfer.ErrAmountInvalid,
		},

		// origin account do not exist
		{
			name: "should return unsuccessful creation of transfer due to not existence of origin accoount",
			fields: fields{
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.ExternalID) (account.Account, error) {
						if ext == id.From(idDestination) {
							return account.Account{
								ID:      2,
								Balance: 2500,
							}, nil
						}
						return account.Account{}, account.ErrNotFound
					},
				},
				tr: &transfer.RepositoryMock{},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					OriginID: id.From(idOrigin),
					DestID:   id.From(idDestination),
					Amount:   2500,
				},
			},
			want:    CreateOutput{},
			wantErr: account.ErrNotFound,
		},

		// destination account do not exist
		{
			name: "should return unsuccessful creation of transfer due to not existence of destination accoount",
			fields: fields{
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.ExternalID) (account.Account, error) {
						if ext != id.From(idDestination) {
							return account.Account{
								ID:      1,
								Balance: 2500,
							}, nil
						}
						return account.Account{}, account.ErrNotFound
					},
				},
				tr: &transfer.RepositoryMock{},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					OriginID: id.From(idOrigin),
					DestID:   id.From(idDestination),
					Amount:   2500,
				},
			},
			want:    CreateOutput{},
			wantErr: account.ErrNotFound,
		},

		// error when beginning transaction
		{
			name: "should return unsuccessful creation of transfer due to error when beginning transaction",
			fields: fields{
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.ExternalID) (account.Account, error) {
						internal := 1
						if ext == id.From(idDestination) {
							internal = 2
						}
						return account.Account{
							ID:      id.ID(internal),
							Balance: 5000,
						}, nil
					},
				},
				tr: &transfer.RepositoryMock{},
				tx: &transaction.RepositoryMock{
					BeginFunc: func(ctx context.Context) (context.Context, error) {
						return ctx, transaction.ErrBeginTransaction
					},
				},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					OriginID: id.From(idOrigin),
					DestID:   id.From(idDestination),
					Amount:   2500,
				},
			},
			want:    CreateOutput{},
			wantErr: transfer.ErrRegistering,
		},

		// error when updating balance of origin
		{
			name: "should return unsuccessful creation of transfer due to error when updating balance of origin",
			fields: fields{
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.ExternalID) (account.Account, error) {
						internal := 1
						if ext == id.From(idDestination) {
							internal = 2
						}
						return account.Account{
							ID:      id.ID(internal),
							Balance: 5000,
						}, nil
					},
					UpdateBalanceFunc: func(ctx context.Context, id id.ExternalID, balance currency.Currency) error {
						return pgx.ErrNoSchema
					},
				},
				tr: &transfer.RepositoryMock{},
				tx: &transaction.RepositoryMock{
					BeginFunc: func(ctx context.Context) (context.Context, error) {
						return ctx, nil
					},
					RollbackFunc: func(ctx context.Context) {},
				},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					OriginID: id.From(idOrigin),
					DestID:   id.From(idDestination),
					Amount:   2500,
				},
			},
			want:    CreateOutput{},
			wantErr: transfer.ErrRegistering,
		},

		// error when updating balance of destination
		{
			name: "should return unsuccessful creation of transfer due to error when updating balance of destination",
			fields: fields{
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.ExternalID) (account.Account, error) {
						internal := 1
						if ext == id.From(idDestination) {
							internal = 2
						}
						return account.Account{
							ID:      id.ID(internal),
							Balance: 5000,
						}, nil
					},
					UpdateBalanceFunc: func(ctx context.Context, ext id.ExternalID, balance currency.Currency) error {
						if ext == id.From(idOrigin) {
							return nil
						}
						return pgx.ErrNoSchema
					},
				},
				tr: &transfer.RepositoryMock{},
				tx: &transaction.RepositoryMock{
					BeginFunc: func(ctx context.Context) (context.Context, error) {
						return ctx, nil
					},
					RollbackFunc: func(ctx context.Context) {},
				},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					OriginID: id.From(idOrigin),
					DestID:   id.From(idDestination),
					Amount:   2500,
				},
			},
			want:    CreateOutput{},
			wantErr: transfer.ErrRegistering,
		},

		// error when creating transfer entity effectively
		{
			name: "should return unsuccessful creation of transfer due to error when creating transfer entity effectively",
			fields: fields{
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.ExternalID) (account.Account, error) {
						internal := 1
						if ext == id.From(idDestination) {
							internal = 2
						}
						return account.Account{
							ID:      id.ID(internal),
							Balance: 5000,
						}, nil
					},
					UpdateBalanceFunc: func(ctx context.Context, ext id.ExternalID, balance currency.Currency) error {
						return nil
					},
				},
				tr: &transfer.RepositoryMock{
					CreateFunc: func(ctx context.Context, tr transfer.Transfer) (transfer.Transfer, error) {
						return tr, pgx.ErrNoSchema
					},
				},
				tx: &transaction.RepositoryMock{
					BeginFunc: func(ctx context.Context) (context.Context, error) {
						return ctx, nil
					},
					RollbackFunc: func(ctx context.Context) {},
				},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					OriginID: id.From(idOrigin),
					DestID:   id.From(idDestination),
					Amount:   2500,
				},
			},
			want:    CreateOutput{},
			wantErr: transfer.ErrRegistering,
		},

		// error when committing transaction
		{
			name: "should return unsuccessful creation of transfer due to error when creating transfer entity effectively",
			fields: fields{
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.ExternalID) (account.Account, error) {
						internal := 1
						if ext == id.From(idDestination) {
							internal = 2
						}
						return account.Account{
							ID:      id.ID(internal),
							Balance: 5000,
						}, nil
					},
					UpdateBalanceFunc: func(ctx context.Context, ext id.ExternalID, balance currency.Currency) error {
						return nil
					},
				},
				tr: &transfer.RepositoryMock{
					CreateFunc: func(ctx context.Context, tr transfer.Transfer) (transfer.Transfer, error) {
						tr.ExternalID = id.From(idTransfer)
						return tr, nil
					},
				},
				tx: &transaction.RepositoryMock{
					CommitFunc: func(ctx context.Context) error {
						return pgx.ErrNoSchema
					},
					BeginFunc: func(ctx context.Context) (context.Context, error) {
						return ctx, nil
					},
					RollbackFunc: func(ctx context.Context) {},
				},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					OriginID: id.From(idOrigin),
					DestID:   id.From(idDestination),
					Amount:   2500,
				},
			},
			want:    CreateOutput{},
			wantErr: transfer.ErrRegistering,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			u := New(test.fields.ac, test.fields.tr, test.fields.tx)
			got, err := u.Create(test.args.ctx, test.args.input)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
