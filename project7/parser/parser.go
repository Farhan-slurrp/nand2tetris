package parser

type CmdType string

const (
	C_ARITHMETIC CmdType = "C_ARITHMETIC"
	C_POP        CmdType = "C_POP"
	C_PUSH       CmdType = "C_PUSH"
	C_LABEL      CmdType = "C_LABEL"
	C_GOTO       CmdType = "C_GOTO"
	C_IF         CmdType = "C_IF"
	C_FUNCTION   CmdType = "C_FUNCTION"
	C_RETURN     CmdType = "C_RETURN"
	C_CALL       CmdType = "C_CALL"
)

type Parser struct{}

type IParser interface {
	HasMoreLines() bool
	Advance()
	commandType() CmdType
	arg1() string
	arg2() int
}

func NewParser() IParser {
	return &Parser{}
}

func (p *Parser) HasMoreLines() bool {
	return false
}

func (p *Parser) Advance() {
	return
}

func (p *Parser) commandType() CmdType {
	return C_ARITHMETIC
}

func (p *Parser) arg1() string {
	return ""
}

func (p *Parser) arg2() int {
	return 0
}
