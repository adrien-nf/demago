package migration

import "fmt"

type TableBuilder struct {
	columns []Column
}

type Column struct {
	Name string
	Type string
}

func (tb *TableBuilder) String(name string) {
	fmt.Println("Creating string column:", name)
	tb.columns = append(tb.columns, Column{Name: name, Type: "string"})
}
