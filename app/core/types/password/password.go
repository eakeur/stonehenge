package password

import (
	"golang.org/x/crypto/bcrypt"
)

// Password is an implementation of a hash and allows you to compare with
// other passwords
type Password string

func From(input string) Password {
	h, _ := bcrypt.GenerateFromPassword([]byte(input), bcrypt.MinCost)
	pwd := string(h)
	return Password(pwd)
}

// Compare validates and compares a password with another one
func (p Password) Compare(input string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(p), []byte(input))
	if err != nil {
		return ErrWrongPassword
	}

	return nil
}

func (p Password) Validate() error {
	if len(string(p)) > 60 {
		return ErrTooBig
	}
	return nil
}
