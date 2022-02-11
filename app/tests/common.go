package tests

func EvaluateDep(specific, generic interface{}) interface{} {
	if specific != nil {
		return specific
	}
	return generic
}
