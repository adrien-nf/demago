package alexander

type Parameters struct {
	Args []string
}

func NewParameters(args []string) Parameters {
	return Parameters{Args: args}
}