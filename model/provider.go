package model

import (
	"context"

	"cloud.google.com/go/firestore"
)

// This type stores, in a normalized way, objects that gives access to data storages
type DataProvider struct {

	// This object allows connection to the database that feeds this application
	Database *firestore.Client

	// This object wraps a single context that can be accessed through the app
	Context context.Context
}
