package alexander

type Command interface {
	Execute(Parameters) error
	Description() string
	Signature() string
}
