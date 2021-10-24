package repository

import (
	model "stonehenge/model"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
)

// This repository controls the IO process between the application and the database for the Account entity
type AccountsRepositoryType struct {

	// Provider
	Provider model.DataProvider
}

// Gets all accounts from the data provider
func (a AccountsRepositoryType) GetAccounts(cpf *string) ([]model.Account, error) {

	collection := a.Provider.Database.Collection("accounts").Query
	if cpf != nil {
		collection = collection.Where("cpf", "==", cpf)
	}

	res, err := collection.Documents(a.Provider.Context).GetAll()
	if err != nil {
		return nil, model.ErrInternal
	}
	accounts := make([]model.Account, len(res))
	for i, doc := range res {
		if doc.Exists() {
			data := doc.Data()
			if data != nil {
				accounts[i] = model.AccountFromMap(data)
			}
		}
	}
	return accounts, nil
}

// Gets the account with the id provided
func (a AccountsRepositoryType) GetAccountById(id string) (*model.Account, error) {
	res, err := a.Provider.Database.Collection("accounts").Where("id", "==", id).Documents(a.Provider.Context).GetAll()
	if err != nil {
		return nil, model.ErrInternal
	}
	if len(res) > 0 {
		doc := res[0]
		if doc.Exists() {
			data := doc.Data()
			if data != nil {
				if data["id"] == id {
					acc := model.AccountFromMap(data)
					return &acc, nil
				}
			}
		}
	}

	return nil, model.ErrNotFound
}

// Adds an account to the database
func (a AccountsRepositoryType) AddAccount(account model.Account) (string, error) {
	account.Id = uuid.New().String()
	account.CreatedAt = time.Now()
	_, _, err := a.Provider.Database.Collection("accounts").Add(a.Provider.Context, account.ToMap())
	if err != nil {
		return "", model.ErrCreating
	}

	return account.Id, nil
}

// Updates an account in the database
func (a AccountsRepositoryType) UpdateAccount(id string, account model.Account) error {
	res, err := a.Provider.Database.Collection("accounts").Where("id", "==", id).Documents(a.Provider.Context).GetAll()
	if err != nil {
		return model.ErrInternal
	}
	if len(res) > 0 {
		doc := res[0]
		if doc.Exists() {
			_, err := doc.Ref.Update(a.Provider.Context, []firestore.Update{
				{Path: "name", Value: account.Name},
				{Path: "cpf", Value: account.Cpf},
				{Path: "secret", Value: account.Secret},
				{Path: "balance", Value: account.Balance},
			})
			if err != nil {
				return model.ErrInternal
			}

			return nil
		}
	}

	return model.ErrNotFound
}

// Removes an account from the database
func (a AccountsRepositoryType) RemoveAccount(id string) error {
	res, err := a.Provider.Database.Collection("accounts").Where("id", "==", id).Documents(a.Provider.Context).GetAll()
	if err != nil {
		return model.ErrInternal
	}
	if len(res) > 0 {
		doc := res[0]
		_, err := doc.Ref.Delete(a.Provider.Context)
		if err != nil {
			return model.ErrInternal
		}
		return nil
	}

	return model.ErrNotFound
}

// Checks if account exists
func (a AccountsRepositoryType) AccountExists(cpf string) (bool, error) {
	res, err := a.Provider.Database.Collection("accounts").Where("cpf", "==", cpf).Documents(a.Provider.Context).GetAll()
	if err != nil {
		return false, model.ErrInternal
	}
	for _, doc := range res {
		return doc.Exists(), nil
	}
	return false, nil
}
