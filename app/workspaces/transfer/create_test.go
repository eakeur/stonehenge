package transfer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	var (
		originID      = id.ExternalFrom("d0052623-0695-4a3a-abf6-887f613dda8e")
		destinationID = id.ExternalFrom("17edb329-4b65-41ba-bb26-5060a1e157ab")
		transferID    = id.ExternalFrom("b8d11928-3eab-45a8-8be3-31411bd120f2")
		unknownID     = id.ExternalFrom("b8d11928-3eab-45a2-8be3-31411bd12a34")
	)

	ctx := context.Background()

	transactions := &transaction.RepositoryMock{BeginResult: ctx}

	accesses := &access.RepositoryMock{
		GetAccessFromContextResult: access.Access{AccountID: originID},
	}

	type args struct {
		ctx   context.Context
		input CreateInput
	}

	type fields struct {
		transactions transaction.Manager
		access       access.Manager
		accounts     account.Repository
		transfers    transfer.Repository
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    CreateOutput
		wantErr error
	}

	tests := []test{

		// Should return CreateOutput with filled fields for successfully created transfer
		{
			name: "return CreateOutput for created transfer",
			fields: fields{
				transfers: &transfer.RepositoryMock{
					CreateFunc: func(ctx context.Context, transfer transfer.Transfer) (transfer.Transfer, error) {
						transfer.ExternalID = transferID
						return transfer, nil
					},
				},
				accounts: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case originID:
							return account.Account{ID: 1, ExternalID: ext, Balance: 5000}, nil
						case destinationID:
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				DestID: destinationID,
				Amount: 2500,
			}},
			want: CreateOutput{RemainingBalance: 2500, TransferId: transferID},
		},

		// Should return ErrNoMoney because origin account had no money to complete the operation
		{
			name: "return ErrNoMoney because origin has not enough budget",
			fields: fields{
				transfers: &transfer.RepositoryMock{},
				accounts: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case originID:
							return account.Account{ID: 1, ExternalID: ext, Balance: 1200}, nil
						case destinationID:
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				DestID: destinationID,
				Amount: 2500,
			}},
			wantErr: account.ErrNoMoney,
		},

		// Should return ErrSameAccount because origin and destination accounts are the same
		{
			name: "return ErrSameAccount because origin equals destination",
			fields: fields{
				transfers: &transfer.RepositoryMock{},
				accounts:  &account.RepositoryMock{},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				DestID: originID,
				Amount: 2500,
			}},
			wantErr: transfer.ErrSameAccount,
		},

		{
			name: "return ErrNoAccessInContext because there is no account logged",
			fields: fields{
				transfers: &transfer.RepositoryMock{},
				accounts:  &account.RepositoryMock{},
				access:    access.RepositoryMock{Error: access.ErrNoAccessInContext},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				DestID: originID,
				Amount: 2500,
			}},
			wantErr: access.ErrNoAccessInContext,
		},

		// Should return ErrAmountInvalid because amount to be transferred is 0
		{
			name: "return ErrAmountInvalid because amount equals zero",
			fields: fields{
				transfers: &transfer.RepositoryMock{},
				accounts:  &account.RepositoryMock{},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				DestID: destinationID,
				Amount: 0,
			}},
			wantErr: transfer.ErrAmountInvalid,
		},

		// Should return ErrNonexistentOrigin because origin account could not be found
		{
			name: "return ErrNonexistentOrigin",
			fields: fields{
				transfers: &transfer.RepositoryMock{},
				access: &access.RepositoryMock{
					CreateResult: access.Access{
						AccountID: unknownID,
					},
				},
				accounts: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case originID:
							return account.Account{ID: 1, ExternalID: ext, Balance: 5000}, nil
						case destinationID:
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				DestID: destinationID,
				Amount: 2500,
			}},
			wantErr: transfer.ErrNonexistentOrigin,
		},

		// Should return ErrNonexistentDestination because destination origin account could not be found
		{
			name: "return ErrNonexistentDestination",
			fields: fields{
				transfers: &transfer.RepositoryMock{},
				accounts: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case originID:
							return account.Account{ID: 1, ExternalID: ext, Balance: 5000}, nil
						case destinationID:
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				DestID: unknownID,
				Amount: 2500,
			}},
			wantErr: transfer.ErrNonexistentDestination,
		},

		// Should return error because origin account's balance could not be updated
		{
			name: "return repository error on updating origin account",
			fields: fields{
				transfers: &transfer.RepositoryMock{},
				accounts: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case originID:
							return account.Account{ID: 1, ExternalID: ext, Balance: 5000}, nil
						case destinationID:
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
					UpdateBalanceFunc: func(ctx context.Context, ext id.External, balance currency.Currency) error {
						switch ext {
						case originID:
							return account.ErrUpdating
						case destinationID:
							return nil
						default:
							return account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				DestID: destinationID,
				Amount: 2500,
			}},
			wantErr: account.ErrUpdating,
		},

		// Should return error because destination account's balance could not be updated
		{
			name: "return repository error on updating destination account",
			fields: fields{
				transfers: &transfer.RepositoryMock{},
				accounts: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case originID:
							return account.Account{ID: 1, ExternalID: ext, Balance: 5000}, nil
						case destinationID:
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
					UpdateBalanceFunc: func(ctx context.Context, ext id.External, balance currency.Currency) error {
						switch ext {
						case originID:
							return nil
						case destinationID:
							return account.ErrUpdating
						default:
							return account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				DestID: destinationID,
				Amount: 2500,
			}},
			wantErr: account.ErrUpdating,
		},

		// Should return error because transfer entity could not be created
		{
			name: "return repository error for creating transfer",
			fields: fields{
				transfers: &transfer.RepositoryMock{Error: transfer.ErrRegistering},
				accounts: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case originID:
							return account.Account{ID: 1, ExternalID: ext, Balance: 5000}, nil
						case destinationID:
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				DestID: destinationID,
				Amount: 2500,
			}},
			wantErr: transfer.ErrRegistering,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tx := test.fields.transactions
			if tx == nil {
				tx = transactions
			}

			tk := test.fields.access
			if tk == nil {
				tk = accesses
			}

			u := New(test.fields.accounts, test.fields.transfers, tx, tk)
			got, err := u.Create(test.args.ctx, test.args.input)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
