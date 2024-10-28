package migration

import "errors"

type Migrator struct {
	migrations []Migration
}

func NewMigrator() *Migrator {
	return &Migrator{}
}

func (m *Migrator) Register(migration Migration) error {
	if migration == nil {
		return errors.New("migration cannot be nil")
	}

	m.migrations = append(m.migrations, migration)
	return nil
}
