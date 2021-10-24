package providers

import (
	"context"
	model "stonehenge/model"
	"stonehenge/repository"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (

	// This repository controls the IO process between the application and the database for the Account entity
	AccountsRepository *repository.AccountsRepositoryType

	// This repository controls the IO process between the application and the database for the Transfer entity
	TransfersRepository *repository.TransfersRepositoryType
)

// Creates an object that allows clients to access the database
func ConnectToDatabase(context context.Context) (*model.DataProvider, error) {
	sa := option.WithCredentialsFile("/go/src/stonehenge/sa.json")
	app, err := firebase.NewApp(context, nil, sa)

	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(context)

	if err != nil {
		return nil, err
	}

	return &model.DataProvider{
		Database: client,
		Context:  context,
	}, nil

}

// Initializes all the repositories and injects the database provider reference
func InjectDependenciesInRepositories(provider *model.DataProvider) {
	AccountsRepository = &repository.AccountsRepositoryType{Provider: *provider}
	TransfersRepository = &repository.TransfersRepositoryType{Provider: *provider}
}
