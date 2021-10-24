package repository

import (
	model "stonehenge/model"
	"time"

	"github.com/google/uuid"
)

// This repository controls the IO process between the application and the database for the Transfer entity.
// As the services implies, transfers can NOT be changed or undone, so there is not a update or delete method in this struct
type TransfersRepositoryType struct {

	// Provider
	Provider model.DataProvider
}

const (
	// Constant that defines that something belongs to the origin account of the transfer
	Origin int = 1

	// Constant that defines that something belongs to the destination account of the transfer
	Destination int = 2
)

// Gets all transfers from the data provider
func (a TransfersRepositoryType) GetTransfers(id string, madeToMe bool) ([]model.Transfer, error) {
	var query string

	// Checks if the applicant wants requests made to them or by them
	if madeToMe {
		query = "account_destination_id"
	} else {
		query = "account_origin_id"

	}
	res, err := a.Provider.Database.Collection("transfers").Where(query, "==", id).Documents(a.Provider.Context).GetAll()
	if err != nil {
		return nil, model.ErrInternal
	}
	transfers := make([]model.Transfer, len(res))
	for i, doc := range res {
		if doc.Exists() {
			data := doc.Data()
			if data != nil {
				transfers[i] = model.TransferFromMap(data)
			}
		}
	}
	return transfers, nil
}

// Gets the transfer with the id provided
func (a TransfersRepositoryType) GetTransferById(id string) (*model.Transfer, error) {
	res, err := a.Provider.Database.Collection("transfers").Where("id", "==", id).Documents(a.Provider.Context).GetAll()
	if err != nil {
		return nil, model.ErrInternal
	}
	if len(res) > 0 {
		doc := res[0]
		if doc.Exists() {
			data := doc.Data()
			if data != nil {
				if data["id"] == id {
					transfer := model.TransferFromMap(data)
					return &transfer, nil
				}
			}
		}
	}

	return nil, model.ErrNotFound
}

// Adds a transfer to the database
func (a TransfersRepositoryType) AddTransfer(transfer model.Transfer) (string, error) {
	transfer.Id = uuid.New().String()
	transfer.CreatedAt = time.Now()
	_, _, err := a.Provider.Database.Collection("transfers").Add(a.Provider.Context, transfer.ToMap())
	if err != nil {
		return "", model.ErrCreating
	}

	return transfer.Id, nil
}
