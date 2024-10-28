package kernel

import (
	"demago/migrations"
	"demago/packages/migration"
)

type MigrationProvider struct{}

func (cp *MigrationProvider) BootstrapMigrations(migrator *migration.Migrator) {
	migrator.Register(&migrations.CreateUsersTable{})
}
