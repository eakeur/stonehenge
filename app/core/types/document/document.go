package document

import (
	"errors"
	"strconv"
	"strings"
)

// Document is an implementation of a string that has other functions to allow validation
type Document string

// Validate validates a document
func (d *Document) Validate() error {
	for _, digit := range *d {
		_, err := strconv.Atoi(string(digit))
		if err != nil {
			return errors.New("")
		}
	}
	if len(*d) != 11 {
		return errors.New("")
	}
	return nil
}

// Normalize removes all special characters from a document string
func Normalize(document string) string {
	return strings.Trim(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(document, ".", ""),
					"-", ""),
				"/", ""),
			",", ""),
		" ")
}
