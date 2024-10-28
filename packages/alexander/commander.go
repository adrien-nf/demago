package alexander

import (
	"errors"
)

type Commander struct {
    commands map[string]Command
}

func NewCommander() *Commander {
    return &Commander{
        commands: make(map[string]Command),
    }
}

func (c *Commander) Register(cmd Command) error {
    if cmd == nil {
        return errors.New("command cannot be nil")
    }
    c.commands[cmd.Signature()] = cmd
    
    return nil
}

func (c *Commander) Execute(commandName string, params Parameters) error {
    cmd, exists := c.commands[commandName]

    if !exists {
        return errors.New("unknown command: " + commandName)
    }

    return cmd.Execute(params)
}