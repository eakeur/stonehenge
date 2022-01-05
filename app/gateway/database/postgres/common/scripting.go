package common

import (
	"database/sql"
	"fmt"
	"strings"
)

// Scanner is meant to standardize both sql.Row and sql.Rows objects that have the Scan method
type Scanner interface {
	Scan(dest ...interface{}) error
}

// Executor  is meant to standardize both sql.tx and sql.db objects that has the Exec method
type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)

	QueryRow(query string, args ...interface{}) *sql.Row

	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func AppendCondition(query string, logic string, condition string) string {
	if strings.Contains(query, " where ") {
		return fmt.Sprintf("%v %v %v", query, logic, condition)
	}
	return fmt.Sprintf("%v where %v", query, condition)
}
