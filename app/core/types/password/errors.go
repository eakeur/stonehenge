package password

import "errors"

var (

	// ErrWrongPassword occurs when a comparison between different passwords happens
	ErrWrongPassword = errors.New("wrong password")
)
