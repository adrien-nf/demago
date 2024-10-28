package migration

type MigrationBuilder struct {
	tables []*TableBuilder
}

func (mb *MigrationBuilder) CreateTable(name string, define func(*TableBuilder)) {
	tb := &TableBuilder{}
	define(tb)
	mb.tables = append(mb.tables, tb)
}

func (mb *MigrationBuilder) DropIfExists(tableName string) {
	//
}