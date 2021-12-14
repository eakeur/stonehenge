package app

import (
	"stonehenge/app/dto"
	m "stonehenge/core/model"
	"stonehenge/core/repositories"
)

type TransferApp interface {

	// Gets all transfers made to or by the specified account in the ID parameter.
	// The toMe bool indicates if the id passed is the destination or the origin account.
	GetAll(id string, toMe bool) ([]dto.TransferDTO, error)

	// Gets the account with the ID specified
	GetById(id string) (*dto.TransferDTO, error)

	// Creates a new transfer
	Add(transfer *m.Transfer) (*string, error)

	// Creates a money transaction between two accounts and creates a new Transfer entity
	Transact(transfer *m.Transfer) (*string, error)
}

type transferApp struct {
	tf repositories.TransferRepository
	ac repositories.AccountRepository
}

func NewTransferApp(transferRepo repositories.TransferRepository, accountRepo repositories.AccountRepository) TransferApp {
	return &transferApp{
		tf: transferRepo,
		ac: accountRepo,
	}
}

// Gets all transfers made to or by the specified account in the ID parameter.
// The toMe bool indicates if the id passed is the destination or the origin account.
func (r *transferApp) GetAll(id string, toMe bool) ([]dto.TransferDTO, error) {
	res, err := r.tf.GetAll(id, toMe)
	if err != nil {
		return nil, err
	}
	transfers := make([]dto.TransferDTO, len(res))
	for i, tf := range res {
		transfers[i] = dto.TransferDTO{
			Id:                   tf.Id,
			AccountOriginId:      tf.AccountOriginId,
			AccountDestinationId: tf.AccountDestinationId,
			Amount:               tf.Amount,
			Date:                 tf.CreatedAt,
		}
	}
	return transfers, nil
}

// Gets the account with the ID specified
func (r *transferApp) GetById(id string) (*dto.TransferDTO, error) {
	res, err := r.tf.GetById(id)
	if err != nil {
		return nil, err
	}

	return &dto.TransferDTO{
		Id:                   res.Id,
		AccountOriginId:      res.AccountOriginId,
		AccountDestinationId: res.AccountDestinationId,
		Amount:               res.Amount,
		Date:                 res.CreatedAt,
	}, nil
}

// Creates a new transfer
func (r *transferApp) Add(transfer *m.Transfer) (*string, error) {
	return r.tf.Add(transfer)
}

// Creates a money transaction between two accounts and creates a new Transfer entity
func (r *transferApp) Transact(transfer *m.Transfer) (*string, error) {
	return r.tf.Transact(transfer)
}

var _ TransferApp = &transferApp{}
