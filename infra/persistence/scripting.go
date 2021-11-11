package persistence

import (
	"database/sql"
	"strings"
)

// Mounts a insert statement based on the values passed
func Insert(ex Executor, table string, values map[string]interface{}) (sql.Result, error) {
	keys, data := decomposeMap(values)
	script := mountInsert(table, keys)
	return ex.Exec(script, data...)
}

func mountInsert(table string, fields []string) string {
	script := "insert into " + table + " ("
	valuesPlaceholder := "("
	for i := 0; i < len(fields); i++ {
		if s := strings.HasSuffix(script, "("); s {
			script += fields[i]
			valuesPlaceholder += "?"
		} else {
			script += ", " + fields[i]
			valuesPlaceholder += ", ?"
		}
	}
	script += ") values " + valuesPlaceholder + ")"
	return script
}

// Mounts an update statement base on the values passed
func Update(ex Executor, table string, values map[string]interface{}, equals map[string]interface{}) (sql.Result, error) {
	keys, data := decomposeMap(values)
	whereKeys, whereData := decomposeMap(equals)
	parameters := append(data, whereData...)
	script := mountUpdate(table, keys, whereKeys)
	return ex.Exec(script, parameters...)
}

func mountUpdate(table string, fields []string, equals []string) string {
	script := "update " + table + " set "
	for i := 0; i < len(fields); i++ {
		if i == 0 {
			script += fields[i] + " = ?"
		} else {
			script += ", " + fields[i] + " = ?"
		}
	}
	if len(equals) > 0 {
		script += " where " + mountWhere(equals)
	}
	return script
}

// Mounts a select statement based on the values passed that returns more than one result
func SelectMany(ex Executor, table string, fields string, equals map[string]interface{}) (*sql.Rows, error) {
	whereKeys, whereData := decomposeMap(equals)
	script := mountSelect(table, fields, whereKeys)
	return ex.Query(script, whereData...)
}

// Mounts a select statement based on the values passed that returns only one result
func SelectOne(ex Executor, table string, fields string, equals map[string]interface{}) *sql.Row {
	whereKeys, whereData := decomposeMap(equals)
	script := mountSelect(table, fields, whereKeys)
	return ex.QueryRow(script, whereData...)
}

func mountSelect(table string, fields string, equals []string) string {
	script := "select " + fields + " from " + table
	if len(equals) > 0 {
		script += " where " + mountWhere(equals)
	}
	return script
}

// Mounts a delete statement based on the values passed
func Delete(ex Executor, table string, equals map[string]interface{}) (sql.Result, error) {
	whereKeys, whereData := decomposeMap(equals)
	script := mountDelete(table, whereKeys)
	return ex.Exec(script, whereData...)
}

func mountDelete(table string, equals []string) string {
	script := "delete from " + table
	if len(equals) > 0 {
		script += " where " + mountWhere(equals)
	}
	return script
}

func decomposeMap(values map[string]interface{}) ([]string, []interface{}) {
	keys := make([]string, 0, len(values))
	data := make([]interface{}, 0, len(values))
	for k, v := range values {
		keys = append(keys, k)
		data = append(data, v)
	}

	return keys, data
}

func mountWhere(fields []string) string {
	script := ""
	for i := 0; i < len(fields); i++ {
		if i == 0 {
			script += fields[i] + " = ?"
		} else {
			script += "and " + fields[i] + " = ?"
		}
	}
	return script
}

// This interface is meant to standardize both sql.Row and sql.Rows objects that has the Scan method
type Scanner interface {
	Scan(dest ...interface{}) error
}

// This interface is meant to standardize both sql.tx and sql.db objects that has the Exec method
type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)

	QueryRow(query string, args ...interface{}) *sql.Row

	Query(query string, args ...interface{}) (*sql.Rows, error)
}
