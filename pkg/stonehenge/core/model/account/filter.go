package account

// Filter stores information that refines the account list, bringing up only what is needed
type Filter struct {

	// Name filters accounts which the holder's name contains this string
	Name string
}
