package transfer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"testing"
)

const (
	originID      string = "d0052623-0695-4a3a-abf6-887f613dda8e"
	destinationID string = "17edb329-4b65-41ba-bb26-5060a1e157ab"
	transferID    string = "b8d11928-3eab-45a8-8be3-31411bd120f2"
	unknownID     string = "b8d11928-3eab-45a2-8be3-31411bd12a34"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	tx := &transaction.RepositoryMock{
		BeginFunc: func(ctx context.Context) (context.Context, error) {
			return ctx, nil
		},
	}

	type args struct {
		ctx   context.Context
		input CreateInput
	}

	type fields struct {
		tx transaction.Transaction
		ac account.Repository
		tr transfer.Repository
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
				tx: tx,
				tr: &transfer.RepositoryMock{
					CreateFunc: func(ctx context.Context, transfer transfer.Transfer) (transfer.Transfer, error) {
						transfer.ExternalID = id.ExternalFrom(transferID)
						return transfer, nil
					},
				},
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case id.ExternalFrom(originID):
							return account.Account{ID: 1, ExternalID: ext, Balance: 5000}, nil
						case id.ExternalFrom(destinationID):
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				OriginID: id.ExternalFrom(originID),
				DestID:   id.ExternalFrom(destinationID),
				Amount:   2500,
			}},
			want: CreateOutput{RemainingBalance: 2500, TransferId: id.ExternalFrom(transferID)},
		},

		// Should return ErrNoMoney because origin account had no money to complete the operation
		{
			name: "return ErrNoMoney because origin has not enough budget",
			fields: fields{
				tx: tx,
				tr: &transfer.RepositoryMock{},
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case id.ExternalFrom(originID):
							return account.Account{ID: 1, ExternalID: ext, Balance: 1200}, nil
						case id.ExternalFrom(destinationID):
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				OriginID: id.ExternalFrom(originID),
				DestID:   id.ExternalFrom(destinationID),
				Amount:   2500,
			}},
			wantErr: account.ErrNoMoney,
		},

		// Should return ErrSameAccount because origin and destination accounts are the same
		{
			name: "return ErrSameAccount because origin equals destination",
			fields: fields{
				tx: tx,
				tr: &transfer.RepositoryMock{},
				ac: &account.RepositoryMock{},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				OriginID: id.ExternalFrom(originID),
				DestID:   id.ExternalFrom(originID),
				Amount:   2500,
			}},
			wantErr: transfer.ErrSameAccount,
		},

		// Should return ErrAmountInvalid because amount to be transferred is 0
		{
			name: "return ErrAmountInvalid because amount equals zero",
			fields: fields{
				tx: tx,
				tr: &transfer.RepositoryMock{},
				ac: &account.RepositoryMock{},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				OriginID: id.ExternalFrom(originID),
				DestID:   id.ExternalFrom(destinationID),
				Amount:   0,
			}},
			wantErr: transfer.ErrAmountInvalid,
		},

		// Should return ErrNonexistentOrigin because origin account could not be found
		{
			name: "return ErrNonexistentOrigin",
			fields: fields{
				tx: tx,
				tr: &transfer.RepositoryMock{},
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case id.ExternalFrom(originID):
							return account.Account{ID: 1, ExternalID: ext, Balance: 5000}, nil
						case id.ExternalFrom(destinationID):
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				OriginID: id.ExternalFrom(unknownID),
				DestID:   id.ExternalFrom(destinationID),
				Amount:   2500,
			}},
			wantErr: transfer.ErrNonexistentOrigin,
		},

		// Should return ErrNonexistentDestination because destination origin account could not be found
		{
			name: "return ErrNonexistentDestination",
			fields: fields{
				tx: tx,
				tr: &transfer.RepositoryMock{},
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case id.ExternalFrom(originID):
							return account.Account{ID: 1, ExternalID: ext, Balance: 5000}, nil
						case id.ExternalFrom(destinationID):
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				OriginID: id.ExternalFrom(originID),
				DestID:   id.ExternalFrom(unknownID),
				Amount:   2500,
			}},
			wantErr: transfer.ErrNonexistentDestination,
		},

		// Should return error because origin account's balance could not be updated
		{
			name: "return repository error on updating origin account",
			fields: fields{
				tx: tx,
				tr: &transfer.RepositoryMock{},
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case id.ExternalFrom(originID):
							return account.Account{ID: 1, ExternalID: ext, Balance: 5000}, nil
						case id.ExternalFrom(destinationID):
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
					UpdateBalanceFunc: func(ctx context.Context, ext id.External, balance currency.Currency) error {
						switch ext {
						case id.ExternalFrom(originID):
							return account.ErrUpdating
						case id.ExternalFrom(destinationID):
							return nil
						default:
							return account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				OriginID: id.ExternalFrom(originID),
				DestID:   id.ExternalFrom(destinationID),
				Amount:   2500,
			}},
			wantErr: account.ErrUpdating,
		},

		// Should return error because destination account's balance could not be updated
		{
			name: "return repository error on updating destination account",
			fields: fields{
				tx: tx,
				tr: &transfer.RepositoryMock{},
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case id.ExternalFrom(originID):
							return account.Account{ID: 1, ExternalID: ext, Balance: 5000}, nil
						case id.ExternalFrom(destinationID):
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
					UpdateBalanceFunc: func(ctx context.Context, ext id.External, balance currency.Currency) error {
						switch ext {
						case id.ExternalFrom(originID):
							return nil
						case id.ExternalFrom(destinationID):
							return account.ErrUpdating
						default:
							return account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				OriginID: id.ExternalFrom(originID),
				DestID:   id.ExternalFrom(destinationID),
				Amount:   2500,
			}},
			wantErr: account.ErrUpdating,
		},

		// Should return error because transfer entity could not be created
		{
			name: "return repository error for creating transfer",
			fields: fields{
				tx: tx,
				tr: &transfer.RepositoryMock{Error: transfer.ErrRegistering},
				ac: &account.RepositoryMock{
					GetByExternalIDFunc: func(ctx context.Context, ext id.External) (account.Account, error) {
						switch ext {
						case id.ExternalFrom(originID):
							return account.Account{ID: 1, ExternalID: ext, Balance: 5000}, nil
						case id.ExternalFrom(destinationID):
							return account.Account{ID: 2, ExternalID: ext, Balance: 5000}, nil
						default:
							return account.Account{}, account.ErrNotFound
						}
					},
				},
			},
			args: args{ctx: context.Background(), input: CreateInput{
				OriginID: id.ExternalFrom(originID),
				DestID:   id.ExternalFrom(destinationID),
				Amount:   2500,
			}},
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
