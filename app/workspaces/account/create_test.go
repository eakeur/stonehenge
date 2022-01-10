package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/types/audits"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func TestAccountCreation(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx   context.Context
		input CreateInput
	}

	type fields struct {
		repo account.Repository
		tx   transaction.Transaction
	}

	type test struct {
		name    string
		fields  fields
		args    args
		want    CreateOutput
		wantErr error
	}
	var accountID = id.New()

	var tests = []test{
		{
			name: "should return successful creation",
			fields: fields{
				repo: &account.RepositoryMock{
					GetWithCPFFunc: func(ctx context.Context, document document.Document) (account.Account, error) {
						return account.Account{}, account.ErrNotFound
					},
					CreateFunc: func(ctx context.Context, account *account.Account) (id.ExternalID, error) {
						return accountID, nil
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
					Document: "97662062015",
					Secret:   password.From("123Qwerty!@#"),
					Name:     "Cesária Évora",
				},
			},
			want: CreateOutput{
				AccountID: accountID,
			},
		},
		{
			name: "should return unsuccessful creation due to already exists error",
			fields: fields{
				repo: &account.RepositoryMock{
					GetWithCPFFunc: func(ctx context.Context, document document.Document) (account.Account, error) {
						return account.Account{
							ID:         1,
							ExternalID: accountID,
							Document:   "97662062015",
							Secret:     password.From("123Qwerty!@#"),
							Name:       "Elza Soares",
							Balance:    500,
							Audit:      audits.Audit{},
						}, nil
					},
				},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					Document: "97662062015",
					Secret:   password.From("123Qwerty!@#"),
					Name:     "Elza Soares",
				},
			},
			want:    CreateOutput{},
			wantErr: account.ErrAlreadyExist,
		},
		{
			name: "should return unsuccessful creation due to invalid document",
			fields: fields{
				repo: &account.RepositoryMock{},
				tx:   &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					Document: "",
					Secret:   password.From("123Qwerty!@#"),
					Name:     "Robin Fenty",
				},
			},
			want:    CreateOutput{},
			wantErr: document.ErrInvalidDocument,
		},
		{
			name: "should return unsuccessful creation due to error on beginning of transaction",
			fields: fields{
				repo: &account.RepositoryMock{
					GetWithCPFFunc: func(ctx context.Context, document document.Document) (account.Account, error) {
						return account.Account{}, nil
					},
				},
				tx: &transaction.RepositoryMock{
					BeginFunc: func(ctx context.Context) (context.Context, error) {
						return ctx, transaction.ErrBeginTransaction
					},
				},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					Document: "97662062015",
					Secret:   password.From("123Qwerty!@#"),
					Name:     "Lina Pereira",
				},
			},
			want:    CreateOutput{},
			wantErr: account.ErrCreating,
		},
		{
			name: "should return unsuccessful creation due to failed commit",
			fields: fields{
				repo: &account.RepositoryMock{
					GetWithCPFFunc: func(ctx context.Context, document document.Document) (account.Account, error) {
						return account.Account{}, account.ErrNotFound
					},
					CreateFunc: func(ctx context.Context, account *account.Account) (id.ExternalID, error) {
						return accountID, nil
					},
				},
				tx: &transaction.RepositoryMock{
					BeginFunc: func(ctx context.Context) (context.Context, error) {
						return ctx, nil
					},
					CommitFunc: func(ctx context.Context) error {
						return pgx.ErrTxCommitRollback
					},
					RollbackFunc: func(ctx context.Context) {},
				},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					Document: "97662062015",
					Secret:   password.From("123Qwerty!@#"),
					Name:     "Cesária Évora",
				},
			},
			want:    CreateOutput{},
			wantErr: account.ErrCreating,
		},
		{
			name: "should return unsuccessful creation due to error on repository",
			fields: fields{
				repo: &account.RepositoryMock{
					GetWithCPFFunc: func(ctx context.Context, document document.Document) (account.Account, error) {
						return account.Account{}, account.ErrNotFound
					},
					CreateFunc: func(ctx context.Context, acc *account.Account) (id.ExternalID, error) {
						return accountID, account.ErrCreating
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
					Document: "97662062015",
					Secret:   password.From("123Qwerty!@#"),
					Name:     "Cesária Évora",
				},
			},
			want:    CreateOutput{},
			wantErr: account.ErrCreating,
		},
		{
			name: "should return unsuccessful creation due to error on checking cpf on database",
			fields: fields{
				repo: &account.RepositoryMock{
					GetWithCPFFunc: func(ctx context.Context, document document.Document) (account.Account, error) {
						return account.Account{}, pgx.ErrNoRows
					},
				},
				tx: &transaction.RepositoryMock{},
			},
			args: args{
				ctx: context.Background(),
				input: CreateInput{
					Document: "97662062015",
					Secret:   password.From("123Qwerty!@#"),
					Name:     "Elza Soares",
				},
			},
			want:    CreateOutput{},
			wantErr: account.ErrCreating,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			u := New(test.fields.repo, test.fields.tx)
			got, err := u.Create(test.args.ctx, test.args.input)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
