package common

import (
	"fmt"
	"strings"
)

const (
	PostgresDuplicateError       = "23505"
	PostgresNonexistentFK        = "23503"
	AccountDocumentUniquenessKey = "accounts_document_key"
	TransferOriginFK             = "transfers_account_origin_id_fkey"
	TransferDestinationFK        = "transfers_account_destination_id_fkey"
)

func AppendCondition(query string, logic string, condition string, paramNumbers ...int) string {
	for _, number := range paramNumbers {
		condition = strings.Replace(condition, "?", fmt.Sprintf("$%v", number), 1)
	}

	if strings.Contains(query, " where ") {
		return fmt.Sprintf("%v %v %v", query, logic, condition)
	}

	return fmt.Sprintf("%v where %v", query, condition)
}
