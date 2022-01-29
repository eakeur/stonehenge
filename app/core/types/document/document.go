package document

import (
	"regexp"
)

var (
	exp = regexp.MustCompile("([^0-9])+")
)

// Document is an implementation of a CPF
// string that has functions to validate CPFs
type Document string

// Validate validates a document
func (d Document) Validate() error {
	if len(d) != 11 {
		return ErrInvalidDocument
	}

	if exp.Match([]byte(d)) {
		return ErrInvalidDocument
	}
	return nil
}

// Normalize removes all special characters from a document string
func (d Document) Normalize() Document {
	return Document(exp.ReplaceAllString(string(d), ""))
}
