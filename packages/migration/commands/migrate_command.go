package commands

import (
	"demago/kernel"
	"demago/packages/alexander"
	"demago/packages/migration"
)

func MigrateCommand() alexander.Command {
	return alexander.NewCommandBuilder().
		WithSignature("migrate").
		WithDescription("Rolls the migrations.").
		WithExecution(func(params alexander.Parameters) error {
		    migrator := migration.NewMigrator()
			
			migrationProvider := &kernel.MigrationProvider{}

			migrationProvider.BootstrapMigrations(migrator)

			migrationPlayer := migration.NewMigrationPlayer(migrator)
			
			if err := migrationPlayer.Apply(); err != nil {
				return err
			}

			return nil
		}).
		Build()
}
