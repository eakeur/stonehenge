package currency

// Currency is an implementation of an int64 with special functions to deal with money
type Currency int64

func New(input int64) Currency {
	return Currency(input)
}

func FromFloat64(input float64) Currency {
	output := int64(input * 100)
	return Currency(output)
}
