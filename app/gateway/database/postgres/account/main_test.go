package account

import (
	"os"
	"stonehenge/app/gateway/database/postgres/postgrestest"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(postgrestest.SetupTest(m))
}
