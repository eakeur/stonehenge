package model

// Login holds useful information about accounts that want to authenticate themselves to the server
type Login struct {
	Identity
}

// Returns a map of this login instance
func (a *Login) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"cpf":    a.Cpf,
		"secret": a.Secret,
	}
}

// Returns an array of this login instance
func (a *Login) ToArray() []interface{} {
	return []interface{}{
		a.Cpf,
		a.Secret,
	}
}

// Returns an instance of login based on the data passed as parameter
func LoginFromMap(mapInput map[string]interface{}) Login {
	return Login{
		Identity: Identity{
			Cpf:    mapInput["cpf"].(string),
			Secret: mapInput["secret"].(string),
		},
	}
}
