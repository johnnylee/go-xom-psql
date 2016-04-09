package xompsql

import (
	"fmt"
	"strings"
)

func (ch *CrudHandler) makeDeleteQuery(table CrudRow) {
	cols := crudColsFromObj(table)

	whereConds := []string{}
	for _, col := range table.KeyCols() {
		whereConds = append(
			whereConds,
			fmt.Sprintf("%v=%v", col, table.ColStoreSql(col)))
	}

	query := fmt.Sprintf("DELETE FROM %v WHERE %v",
		table.Table(),
		strings.Join(whereConds, " AND "))

	ch.deleteQuery, ch.deleteCols = insertPlaceholders(query, cols)
}

func (ch *CrudHandler) Delete(dbrw DBRW, obj interface{}) error {
	args := ch.makeArgs(obj, ch.deleteCols)
	_, err := dbrw.Exec(ch.deleteQuery, args...)
	return err
}
