package password

import (
	"crypto/md5"
	"encoding/hex"
)

// Password is an implementation of a string that automatically hashes its content and allows you to compare with
// other strings
type Password string

// CompareWithString compares a string input with a hashed string
func (p *Password) CompareWithString(input string) error {
	if *p != New(input) {
		return ErrWrongPassword
	}
	return nil
}

// New creates a hashed password from a string
func New(input string) Password {
	return Password(hash(input))
}

// hash receives a string and returns a hash of it
func hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
