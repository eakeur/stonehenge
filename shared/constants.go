package shared

const JWT_KEY string = "EDAF12D5D997C58B1962FD8350E8B1C158447B5D1002DABA4F551BC3CD38F236"

const TOKEN_VALID_MINUTES = 10

type key int

const (
	ContextAccount key = iota
)
