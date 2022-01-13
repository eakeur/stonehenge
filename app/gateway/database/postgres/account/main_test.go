package account

import (
	"os"
	"stonehenge/app/gateway/database/postgres/postgres_test"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(postgres_test.SetupTest(m))
}
