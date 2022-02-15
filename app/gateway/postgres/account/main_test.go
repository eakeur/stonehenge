package account

import (
	"os"
	"stonehenge/app/gateway/postgres/tests"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(tests.SetupTest(m))
}
