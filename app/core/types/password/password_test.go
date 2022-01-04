package password_test

import (
	"stonehenge/app/core/types/password"
	"testing"
)

const (
	example = "Awd@@Awd2022"
	wrong   = "Aws@@Aws2021"
)

var pass = password.From(example)

func TestSuccessfulPasswordComparison(t *testing.T) {
	err := pass.Compare(example)
	if err != nil {
		t.Error("Expected successful password comparison")
	}

}

func TestFailingPasswordComparison(t *testing.T) {
	err := pass.Compare(wrong)
	if err == nil {
		t.Error("Expected password comparison to throw an ErrWrongPassword")
	}

}
