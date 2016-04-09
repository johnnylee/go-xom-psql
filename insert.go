package xompsql

import (
	"fmt"
	"strings"
)

func (ch *CrudHandler) makeInsertQuery(table CrudRow) {
	cols := crudColsFromObj(table)

	colNames := []string{}
	colVals := []string{}

	for _, col := range cols {
		colNames = append(colNames, col.Name)
		colVals = append(colVals, table.ColStoreSql(col.Name))
	}

	query := fmt.Sprintf("INSERT INTO %v(%v) VALUES(%v)",
		table.Table(),
		strings.Join(colNames, ","),
		strings.Join(colVals, ","))

	ch.insertQuery, ch.insertCols = insertPlaceholders(query, cols)
}

func (ch *CrudHandler) Insert(dbrw DBRW, obj interface{}) error {
	args := ch.makeArgs(obj, ch.insertCols)
	_, err := dbrw.Exec(ch.insertQuery, args...)
	return err
}
