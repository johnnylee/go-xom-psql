package xompsql

import (
	"database/sql"
	"reflect"
)

type DBRW interface {
	Exec(string, ...interface{}) (sql.Result, error)
}

type DBRO interface {
	Get(interface{}, string, ...interface{}) error
	Select(interface{}, string, ...interface{}) error
}

type CrudRow interface {
	Table() string
	KeyCols() []string // Update or Get.

	ListWhereCols() []string
	UpdateExcludeCols() []string

	ColStoreSql(string) string
	ColLoadSql(string) string
}

type crudCol struct {
	Name       string
	FieldIndex int
}

func crudColsFromObj(table CrudRow) []crudCol {
	//obj := reflect.ValueOf(table).Interface()
	objType := reflect.TypeOf(table)

	cols := []crudCol{}
	for i := 0; i < objType.NumField(); i++ {
		name := objType.Field(i).Tag.Get("db")
		if len(name) == 0 {
			continue
		}
		col := crudCol{name, i}
		cols = append(cols, col)
	}
	return cols
}

type CrudHandler struct {
	insertQuery string
	insertCols  []int // List of field indices per column.

	getQuery string
	getCols  []int

	listQuery string
	listCols  []int

	updateQuery string
	updateCols  []int

	deleteQuery string
	deleteCols  []int
}

func NewCrudHandler(table CrudRow) *CrudHandler {
	ch := CrudHandler{}
	ch.makeInsertQuery(table)
	ch.makeUpdateQuery(table)
	ch.makeDeleteQuery(table)
	ch.makeGetQuery(table)
	ch.makeListQuery(table)
	return &ch
}

func (ch *CrudHandler) makeArgs(obj interface{}, cols []int) []interface{} {
	args := make([]interface{}, len(cols))

	objVal := reflect.ValueOf(obj)
	for i, fieldIdx := range cols {
		args[i] = objVal.Field(fieldIdx).Interface()
	}
	return args
}
