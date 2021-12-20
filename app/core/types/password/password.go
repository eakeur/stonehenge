package password

import (
	"crypto/md5"
	"encoding/hex"
)

// Password is an implementation of a string that automatically hashes its content and allows you to compare with
// other strings
type Password string

// CompareWithString validates and compares a password with another one
func (p *Password) CompareWithString(input Password) error {
	if *p != Password(input) {
		return ErrWrongPassword
	}
	return nil
}

// Hash receives a string and returns a Hash of it
func (p *Password) Hash() Password {
	hash := md5.Sum([]byte(*p))
	return Password(hex.EncodeToString(hash[:]))
}
