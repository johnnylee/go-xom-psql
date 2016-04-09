package xompsql

import (
	"fmt"
	"strings"
)

func insertPlaceholders(query string, cols []crudCol) (string, []int) {
	argCols := []int{}

	// Find columns still in the query, and replace placeholders.
	colIdx := 0

	for _, col := range cols {
		if !strings.Contains(query, ":"+col.Name+":") {
			continue
		}

		colIdx += 1
		argCols = append(argCols, col.FieldIndex)

		query = strings.Replace(
			query, ":"+col.Name+":", fmt.Sprintf("$%v", colIdx), -1)
	}

	return query, argCols
}

func stringInSlice(s string, a []string) bool {
	for _, x := range a {
		if x == s {
			return true
		}
	}
	return false
}
