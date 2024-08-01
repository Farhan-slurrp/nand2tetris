package parser

type Parser struct {
	filename            string
	current_instruction *string
}

type IParser interface {
	hasMoreLines() bool
	advance()
	instructionType() string
	symbol() string
	dest() string
	comp() string
	jump() string
}
