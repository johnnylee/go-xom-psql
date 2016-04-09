package xompsql

import (
	"fmt"
	"strings"
)

func (ch *CrudHandler) makeListQuery(table CrudRow) {
	cols := crudColsFromObj(table)

	colNames := []string{}
	for _, col := range cols {
		colNames = append(colNames, table.ColLoadSql(col.Name))
	}

	whereConds := []string{}
	for _, col := range table.ListWhereCols() {
		whereConds = append(
			whereConds,
			fmt.Sprintf("%v=%v", col, table.ColStoreSql(col)))
	}

	query := fmt.Sprintf("SELECT %v FROM %v WHERE %v",
		strings.Join(colNames, ","),
		table.Table(),
		strings.Join(whereConds, " AND "))

	ch.listQuery, ch.listCols = insertPlaceholders(query, cols)
}

func (ch *CrudHandler) List(
	dbro DBRO,
	obj, objs interface{},
	orderBy []string, orderDir string,
	limit, offset int) error {

	query := ch.listQuery
	orderClause := ""

	if len(orderBy) != 0 {
		orderClause = fmt.Sprintf("ORDER BY %v %v",
			strings.Join(orderBy, ","),
			orderDir)
	}

	query = fmt.Sprintf("%v %v LIMIT %v OFFSET %v",
		query, orderClause, limit, offset)

	args := ch.makeArgs(obj, ch.listCols)
	return dbro.Select(objs, query, args...)
}
