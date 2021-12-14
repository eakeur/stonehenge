package model

import (
	"testing"
)

func TestTransfer(t *testing.T) {
	origin := Account{
		Id:      "1",
		Balance: 500,
	}

	dest := Account{
		Id:      "2",
		Balance: 500,
	}

	transfer := Transfer{AccountOriginId: origin.Id, AccountDestinationId: dest.Id, Amount: 300}

	err := transfer.TransferMoney(&origin, &dest)

	if err != nil {
		t.Error("expected being allowed to transfer $3,00")
	}

}

func TestSameAccountTransfer(t *testing.T) {
	origin := Account{
		Id:      "1",
		Balance: 500,
	}

	dest := Account{
		Id:      "1",
		Balance: 500,
	}

	transfer := Transfer{AccountOriginId: origin.Id, AccountDestinationId: dest.Id, Amount: 300}

	err := transfer.TransferMoney(&origin, &dest)

	if err == nil {
		t.Error("expected not being allowed to transfer between same accounts")
	}

}

func TestTransferWithNoBudgetApplicant(t *testing.T) {
	origin := Account{
		Id:      "1",
		Balance: 200,
	}

	dest := Account{
		Id:      "2",
		Balance: 500,
	}

	transfer := Transfer{AccountOriginId: origin.Id, AccountDestinationId: dest.Id, Amount: 300}

	err := transfer.TransferMoney(&origin, &dest)

	if err == nil {
		t.Error("expected not being allowed to transfer from account with not enough balance")
	}

}

func TestTransferWithNoAmount(t *testing.T) {
	origin := Account{
		Id:      "1",
		Balance: 500,
	}

	dest := Account{
		Id:      "2",
		Balance: 500,
	}

	transfer := Transfer{AccountOriginId: origin.Id, AccountDestinationId: dest.Id, Amount: 0}

	err := transfer.TransferMoney(&origin, &dest)

	if err == nil {
		t.Error("expected not being allowed to transfer no quantity of money")
	}

}

func TestTransferWithNegativeAmount(t *testing.T) {
	origin := Account{
		Id:      "1",
		Balance: 500,
	}

	dest := Account{
		Id:      "2",
		Balance: 500,
	}

	transfer := Transfer{AccountOriginId: origin.Id, AccountDestinationId: dest.Id, Amount: -250}

	err := transfer.TransferMoney(&origin, &dest)

	if err == nil {
		t.Error("expected not being allowed to transfer negative quantity of money")
	}

}
