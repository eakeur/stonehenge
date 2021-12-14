package security_test

import (
	"stonehenge/infra/security"
	"testing"

	"github.com/google/uuid"
)

const EXIPIRED_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY2NzMxMDQsIkFjY291bnRJZCI6ImQ5ZGZmN2U4LThhNDgtNGUzOS04Njk1LWJjZGNlMTdhYjM0ZSJ9.vzCx9Gkv1FGuPrhpidev4tIXKFXT82vckYGcBuZrbzg"

func TestTokenAPI(t *testing.T) {

	id := uuid.New().String()
	token, err := security.CreateToken(id)
	if err != nil {
		t.Error("expected being allowed to create token for id " + id)
	}

	details, err := security.ExtractToken(token)
	if err != nil {
		t.Error("expected being allowed to extract token with id " + id)
	}

	if id != *details.AccountId {
		t.Error("expected consistency between entity id and token id")
	}

	_, err = security.ExtractToken("faketoken")
	if err == nil {
		t.Error("expected not being allowed to extract fake token")
	}

	_, err = security.ExtractToken(EXIPIRED_TOKEN)
	if err == nil {
		t.Error("expected not being allowed to extract expired token")
	}

}
