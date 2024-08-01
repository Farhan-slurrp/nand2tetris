package parser

type Parser struct {
	file               [][]byte
	currentInstruction int
	currentLine        string
}

type IParser interface {
	HasMoreLines() bool
	Advance()
	InstructionType() string
	Symbol() string
	Dest() string
	Comp() string
	Jump() string
}
