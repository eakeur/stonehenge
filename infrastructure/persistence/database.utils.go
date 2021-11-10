package persistence

import (
	"strings"
)

// Mounts a insert statement based on the values passed
func MountInsert(table string, values map[string]interface{}) (string, []interface{}) {
	script := "insert into " + table + " ("
	valuesArray := make([]interface{}, 0, len(values))
	for k, v := range values {
		if s := strings.HasSuffix(script, "("); s {
			script += k
		} else {
			script += ", " + k
		}
		valuesArray = append(valuesArray, v)
	}
	script += ") values ("

	for i := 0; i < len(values); i++ {
		if s := strings.HasSuffix(script, "("); s {
			script += "$" + string(rune(i+1))
		} else {
			script += ", " + "$" + string(rune(i+1))
		}

	}
	script += ")"

	return script, valuesArray
}

// Mounts an update statement base on the values passed
func MountUpdate(table string, values map[string]interface{}, equals map[string]interface{}) (string, []interface{}) {
	script := "update " + table + " set "
	valuesArray := make([]interface{}, 0, len(values))
	index := 1
	for k, v := range values {
		if s := strings.HasSuffix(script, "set "); s {
			script += k + " = " + "$" + string(rune(index))
		} else {
			script += ", " + k + " = " + "$" + string(rune(index))
		}
		valuesArray = append(valuesArray, v)
		index++
	}
	if equals != nil {
		script += " where "
	}

	for k, v := range equals {
		if s := strings.HasSuffix(script, "where "); s {
			script += k + " = ? "
		} else {
			script += "and " + k + " = ? "
		}
		valuesArray = append(valuesArray, v)
	}

	return script, valuesArray
}

// Mounts a select statement based on the values passed
func MountSelect(table string, fields string, equals map[string]interface{}) (string, []interface{}) {
	script := "select " + fields + " from " + table
	if equals != nil {
		script += " where "
	}

	valuesArray := make([]interface{}, 0, len(equals))
	for k, v := range equals {
		if s := strings.HasSuffix(script, "where "); s {
			script += k + " = ? "
		} else {
			script += "and " + k + " = ? "
		}
		valuesArray = append(valuesArray, v)
	}
	return script, valuesArray
}

// Mounts a delete statement based on the values passed
func MountDelete(table string, equals map[string]interface{}) (string, []interface{}) {
	script := "delete from " + table + " set "
	valuesArray := make([]interface{}, 0, len(equals))
	if equals != nil {
		script += " where "
	}

	for k, v := range equals {
		if s := strings.HasSuffix(script, "where "); s {
			script += k + " = ? "
		} else {
			script += "and " + k + " = ? "
		}
		valuesArray = append(valuesArray, v)
	}

	return script, valuesArray
}

// This interface is meant to standardize both sql.Row and sql.Rows objects that has the Scan method
type Scanner interface {
	Scan(dest ...interface{}) error
}
