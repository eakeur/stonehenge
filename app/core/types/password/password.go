package password

import (
	"golang.org/x/crypto/bcrypt"
)

// Password is an implementation of a hash and allows you to compare with
// other passwords
type Password []byte

func From(input string) Password {
	h, _ := bcrypt.GenerateFromPassword([]byte(input), bcrypt.MinCost)
	return h
}

// Compare validates and compares a password with another one
func (p Password) Compare(input string) error {
	err := bcrypt.CompareHashAndPassword(p, []byte(input))
	if err != nil {
		return ErrWrongPassword
	}

	return nil
}
