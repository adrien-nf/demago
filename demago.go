package main

import (
	"demago/kernel"
	"demago/packages/alexander"
	bootstrap_commands "demago/packages/bootstrap/commands"
	migration_commands "demago/packages/migration/commands"
	"fmt"
	"os"
)

func main() {
    if len(os.Args) < 2 {
        return
    }

	commandsProvider := &kernel.CommandProvider{}
	commander := alexander.NewCommander()
	
	commander.Register(bootstrap_commands.ServeCommand())
	commander.Register(migration_commands.MigrateCommand())

	commandName := os.Args[1]
    args := os.Args[2:]
    params := alexander.NewParameters(args)

    if err := commander.Execute(commandName, params); err != nil {
        fmt.Println("Error:", err)
    }
	

	commandsProvider.BootstrapCommands(commander)
}