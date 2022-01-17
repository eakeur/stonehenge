package account

import (
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"time"
)

type AuthenticationRequest struct {
	Document document.Document
	Secret   string
}

type CreateInput struct {
	Document document.Document
	Secret   password.Password
	Name     string
}

type CreateOutput struct {
	AccountID id.External
	CreatedAt time.Time
}

type GetBalanceResponse struct {
	Balance currency.Currency
}

type Reference struct {
	ExternalID id.External
	Name       string
}
