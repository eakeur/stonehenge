package model

// Login holds useful information about accounts that want to authenticate themselves to the server
type Login struct {

	// The unique document that represents the owner of the account
	Cpf string `json:"cpf"`

	// A string password defined by the owner
	Secret string `json:"secret"`
}

// Returns a map of this login instance
func (a *Login) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"cpf":    a.Cpf,
		"secret": a.Secret,
	}
}

// Returns an instance of login based on the data passed as parameter
func LoginFromMap(mapInput map[string]interface{}) Login {
	return Login{
		Cpf:    mapInput["cpf"].(string),
		Secret: mapInput["secret"].(string),
	}
}
