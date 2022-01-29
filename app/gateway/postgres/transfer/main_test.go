package transfer

import (
	"os"
	"stonehenge/app/gateway/postgres/postgrestest"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(postgrestest.SetupTest(m))
}
