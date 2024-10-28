package alexander

type CommandBuilder struct {
	name        string
	signature   string
	description string
	execution   func(Parameters) error
}

func NewCommandBuilder() *CommandBuilder {
	return &CommandBuilder{}
}

func (cb *CommandBuilder) WithSignature(signature string) *CommandBuilder {
	cb.signature = signature
	return cb
}

func (cb *CommandBuilder) WithDescription(description string) *CommandBuilder {
	cb.description = description
	return cb
}

func (cb *CommandBuilder) WithExecution(execution func(Parameters) error) *CommandBuilder {
	cb.execution = execution
	return cb
}

func (cb *CommandBuilder) Build() Command {
	return &builtCommand{
		name:        cb.name,
		signature:   cb.signature,
		description: cb.description,
		execution:   cb.execution,
	}
}

type builtCommand struct {
	name        string
	signature   string
	description string
	execution   func(Parameters) error
}

func (bc *builtCommand) Execute(params Parameters) error {
	return bc.execution(params)
}

func (bc *builtCommand) Description() string {
	return bc.description
}

func (bc *builtCommand) Signature() string {
	return bc.signature
}
