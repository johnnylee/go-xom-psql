package xompsql

import (
	"fmt"
	"reflect"
	"strings"
)

func (ch *CrudHandler) makeGetQuery(table CrudRow) {
	cols := crudColsFromObj(table)

	colNames := []string{}
	for _, col := range cols {
		colNames = append(colNames, table.ColLoadSql(col.Name))
	}

	whereConds := []string{}
	for _, col := range table.KeyCols() {
		whereConds = append(
			whereConds,
			fmt.Sprintf("%v=%v", col, table.ColStoreSql(col)))
	}

	query := fmt.Sprintf("SELECT %v FROM %v WHERE %v",
		strings.Join(colNames, ","),
		table.Table(),
		strings.Join(whereConds, " AND "))

	ch.getQuery, ch.getCols = insertPlaceholders(query, cols)
}

func (ch *CrudHandler) Get(dbro DBRO, objPtr interface{}) error {
	obj := reflect.ValueOf(objPtr).Elem().Interface()
	args := ch.makeArgs(obj, ch.getCols)
	return dbro.Get(objPtr, ch.getQuery, args...)
}
