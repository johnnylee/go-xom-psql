package xompsql

import (
	"fmt"
	"strings"
)

func (ch *CrudHandler) makeUpdateQuery(table CrudRow) {
	cols := crudColsFromObj(table)

	updateCols := []string{}
	for _, col := range cols {
		if stringInSlice(col.Name, table.KeyCols()) {
			continue
		}
		if stringInSlice(col.Name, table.UpdateExcludeCols()) {
			continue
		}
		updateCols = append(
			updateCols,
			fmt.Sprintf("%v=%v",
				col.Name, table.ColStoreSql(col.Name)))
	}

	whereConds := []string{}
	for _, col := range table.KeyCols() {
		whereConds = append(
			whereConds,
			fmt.Sprintf("%v=%v", col, table.ColStoreSql(col)))
	}

	query := fmt.Sprintf("UPDATE %v SET %v WHERE %v",
		table.Table(),
		strings.Join(updateCols, ","),
		strings.Join(whereConds, " AND "))

	ch.updateQuery, ch.updateCols = insertPlaceholders(query, cols)
}

func (ch *CrudHandler) Update(dbrw DBRW, obj interface{}) error {
	args := ch.makeArgs(obj, ch.updateCols)
	_, err := dbrw.Exec(ch.updateQuery, args...)
	return err
}
