package password_test

import (
	"github.com/stretchr/testify/assert"
	"stonehenge/app/core/types/password"
	"testing"
)

const (
	example = "Awd@@Awd2022"
	wrong   = "Aws@@Aws2021"
)

var (
	pass  = password.From(example)
	cases = []TestCase{
		{
			"matching passwords",
			example,
			nil,
		},
		{
			"different passwords",
			wrong,
			password.ErrWrongPassword,
		},
	}
)

type TestCase struct {
	description    string
	passwordToTest string
	expectedError  error
}

func TestPasswordComparison(t *testing.T) {
	for _, cs := range cases {
		t.Run(cs.description, func(t *testing.T) {
			err := pass.Compare(cs.passwordToTest)
			assert.Equal(t, cs.expectedError, err)
		})
	}
}


