package currency

// Currency is an implementation of an int64 with special functions to deal with money
type Currency int64

func (c Currency) ToStandardCurrency() float64 {
	return float64(c) / 100
}

func FromStandardCurrency(input float64) Currency {
	output := int64(input * 100)
	return Currency(output)
}


