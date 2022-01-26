package authentication

import (
	"stonehenge/app/core/types/document"
)

type AuthenticationRequest struct {
	Document document.Document
	Secret   string
}
