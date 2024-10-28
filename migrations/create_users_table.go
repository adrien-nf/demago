package migrations

import (
	"demago/packages/migration"
)

type CreateUsersTable struct{}

func (m *CreateUsersTable) Up(builder *migration.MigrationBuilder) {
    builder.CreateTable("users", func(table *migration.TableBuilder) {
        table.String("name")
        table.String("email")
        table.String("password")
    })
}

func (m *CreateUsersTable) Down(builder *migration.MigrationBuilder) {
	builder.DropIfExists("users")
}

var _ migration.Migration = (*CreateUsersTable)(nil)
