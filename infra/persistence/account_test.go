package persistence_test

import (
	"stonehenge/infra/persistence"
	"testing"
)

func TestTokenAPI(t *testing.T) {
	_, err := persistence.NewWorkspace("", "", "", "")
	if err != nil {
		return
	}
}
