package main

import (
	"demago/kernel"
	"demago/packages/alexander"
	"demago/packages/bootstrap/commands"
	"fmt"
	"os"
)

func main() {
    if len(os.Args) < 2 {
        return
    }

	commandsProvider := &kernel.CommandProvider{}
	commander := alexander.NewCommander()
	
	commander.Register(commands.ServeCommand())

	commandName := os.Args[1]
    args := os.Args[2:]
    params := alexander.NewParameters(args)

    if err := commander.Execute(commandName, params); err != nil {
        fmt.Println("Error:", err)
    }
	

	commandsProvider.BootstrapCommands(commander)
}