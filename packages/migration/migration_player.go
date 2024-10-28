package migration

import (
	"fmt"
	"reflect"
)

type MigrationPlayer struct {
    migrator *Migrator
}

func NewMigrationPlayer(migrator *Migrator) *MigrationPlayer {
    return &MigrationPlayer{migrator: migrator}
}

func getMigrationName(migration Migration) string {
    return reflect.TypeOf(migration).Elem().Name()
}

func (mp *MigrationPlayer) Apply() error {
    for _, migration := range mp.migrator.migrations {
        builder := &MigrationBuilder{}

        fmt.Println("Applying migration:", getMigrationName(migration))

        migration.Up(builder)

        fmt.Println("Succesfully migrated:", getMigrationName(migration))
    }
    return nil
}

func (mp *MigrationPlayer) Revert() error {
    for _, migration := range mp.migrator.migrations {
        builder := &MigrationBuilder{}
        fmt.Println("Reverting migration:", getMigrationName(migration))

        migration.Down(builder)

        fmt.Println("Succesfully reverted:", getMigrationName(migration))
    }
    return nil
}
