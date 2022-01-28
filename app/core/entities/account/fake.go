package account

import "stonehenge/app/core/types/password"

var uniquePasswordInstance = password.From("12345678")

func GetFakeAccounts() []Account {
	return []Account{
		{
			Document: "70830052062",
			Secret:   password.From("12345678"),
			Name:     "John Reis",
			Balance:  105000,
		},
		{
			Document: "24388516007",
			Secret:   password.From("12345678"),
			Name:     "Wagner Reis",
			Balance:  450000,
		},
		{
			Document: "05161964057",
			Secret:   password.From("12345678"),
			Name:     "Spencer Reis",
			Balance:  502200,
		},
		{
			Document: "24788516002",
			Secret:   password.From("12345678"),
			Name:     "Lina Pereira",
			Balance:  994500,
		},
		{
			Document: "24385516005",
			Secret:   password.From("12345678"),
			Name:     "Elza Soares",
			Balance:  488500,
		},
		{
			Document: "24384516008",
			Secret:   password.From("12345678"),
			Name:     "Jur Arras",
			Balance:  1004500,
		},
	}
}

func GetFakeAccount() Account {
	return Account{
		Document: "24788516002",
		Secret:   uniquePasswordInstance,
		Name:     "Lina Pereira",
		Balance:  50000000,
	}
}
