package model

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

// An object with properties that identifies an user
type Identity struct {

	// The unique document that represents the user
	Cpf string `json:"cpf"`

	// A string password defined by the owner
	Secret string `json:"secret"`
}

func (i *Identity) NormalizeCPF() {
	i.Cpf = strings.Trim(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(i.Cpf, ".", ""), "-", ""), "/", ""), ",", ""), " ")
}

func (i *Identity) IsCPFValid() bool {
	i.NormalizeCPF()
	for _, digit := range i.Cpf {
		_, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}
	}
	return len(i.Cpf) == 11
}

// Creates a hash of the secret property
func (i *Identity) HashSecret() {
	i.Secret = i.hash(i.Secret)
}

func (i *Identity) hash(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

// Verifies if the password passed as a parameter is equal to the hash of the actual one
func (i *Identity) ValidatePass(try string) error {
	if i.Secret != i.hash(try) {
		return ErrWrongPassword
	}
	return nil
}

// Verifies if the hash passed as a parameter is equal to the hash of the actual one
func (i *Identity) ValidateHash(try string) error {
	if i.Secret != try {
		return ErrWrongPassword
	}
	return nil
}
