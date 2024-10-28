package commands

import (
	"demago/packages/alexander"
	app "demago/packages/bootstrap"
)

func ServeCommand() alexander.Command {
	return alexander.NewCommandBuilder().
		WithSignature("serve").
		WithDescription("Runs the server.").
		WithExecution(func(params alexander.Parameters) error {
		    appInstance := &app.App{}

    		appInstance.Start()

			return nil;
		}).
		Build()
}
