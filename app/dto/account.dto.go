package dto

type SafeAccountDTO struct {
	// The unique identifier of this entity
	Id string `json:"id"`

	// The name of the account's owner
	Name string `json:"name"`
}

// Account holds useful information about accounts
type AccountDTO struct {
	// The unique identifier of this entity
	Id string `json:"id"`

	// The name of the account's owner
	Name string `json:"name"`

	// The unique document that represents the owner of this account
	Cpf string `json:"cpf"`

	// A hash of the password defined by the owner
	Secret string `json:"secret"`

	// The actual balance of this account
	Balance int64 `json:"balance"`
}
