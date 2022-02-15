package tests

import (
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/password"
)

var uniquePasswordInstance = password.From("12345678")

func GetFakeAccounts() []account.Account {
	return []account.Account{
		{
			Document: "70830052062",
			Secret:   password.From("12345678"),
			Name:     "John Reis",
			Balance:  2500,
		},
		{
			Document: "24388516007",
			Secret:   password.From("12345678"),
			Name:     "Wagner Reis",
			Balance:  4500,
		},
		{
			Document: "05161964057",
			Secret:   password.From("12345678"),
			Name:     "Spencer Reis",
			Balance:  5000,
		},
		{
			Document: "24788516002",
			Secret:   password.From("12345678"),
			Name:     "Lina Pereira",
			Balance:  4500,
		},
		{
			Document: "24385516005",
			Secret:   password.From("12345678"),
			Name:     "Elza Soares",
			Balance:  4500,
		},
		{
			Document: "24384516008",
			Secret:   password.From("12345678"),
			Name:     "Jur Arras",
			Balance:  4500,
		},
	}
}

func GetFakeAccount() account.Account {
	return account.Account{
		Document: "24788516002",
		Secret:   uniquePasswordInstance,
		Name:     "Lina Pereira",
		Balance:  4500,
	}
}

func GetFakeTransfers() []transfer.Transfer {
	return []transfer.Transfer{
		{
			OriginID:      1,
			DestinationID: 2,
			Amount:        500,
		},
		{
			OriginID:      1,
			DestinationID: 2,
			Amount:        500,
		},
		{
			OriginID:      2,
			DestinationID: 1,
			Amount:        500,
		},
		{
			OriginID:      2,
			DestinationID: 1,
			Amount:        500,
		},
		{
			OriginID:      1,
			DestinationID: 2,
			Amount:        500,
		},
	}
}
