package code

type Code struct {
	destInstruction string
	compInstruction string
	jumpInstruction string
}

type ICode interface {
	Dest() string
	Comp() (string, string)
	Jump() string
}
