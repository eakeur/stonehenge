package shared

// This function validates if the CPF document is valid. On purpose it validates only its length, so that you, the tester, don't need to provide an
// existing CPF
func IsCPFValid(cpf string) bool {
	return len(cpf) == 11
}
