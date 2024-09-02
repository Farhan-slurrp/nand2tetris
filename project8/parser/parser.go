package parser

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

type Parser struct {
	currentCommad int
	file          [][]byte
}

type IParser interface {
	HasMoreLines() bool
	Advance() []byte
	CommandType() CmdType
	GetArg1() string
	GetArg2() int
	GetCurrentLine() string
}

func NewParser(filename string) IParser {
	file, err := os.ReadFile(fmt.Sprintf("./vm/%s.vm", filename))
	if err != nil {
		fmt.Println("Can't read file:", fmt.Sprintf("./vm/%s.vm", filename))
		panic(err)
	}

	array := bytes.Split(file, []byte("\n"))
	return &Parser{
		currentCommad: 0,
		file:          array,
	}
}

func (p *Parser) HasMoreLines() bool {
	return p.currentCommad < len(p.file)
}

func (p *Parser) Advance() []byte {
	curr := p.currentCommad
	p.currentCommad += 1
	return p.file[curr]
}

func (p *Parser) CommandType() CmdType {
	arg0 := strings.Split(p.GetCurrentLine(), " ")[0]

	switch arg0 {
	case "add", "sub", "eq", "gt", "lt", "and", "or", "neg", "not":
		return C_ARITHMETIC
	case "push":
		return C_PUSH
	case "pop":
		return C_POP
	case "label":
		return C_LABEL
	case "goto":
		return C_GOTO
	case "if-goto":
		return C_IF
	case "function":
		return C_FUNCTION
	case "return":
		return C_RETURN
	case "call":
		return C_CALL
	default:
		panic("cannot get the command type")
	}
}

func (p *Parser) GetArg1() string {
	commandType := p.CommandType()
	if commandType == C_RETURN {
		panic("cannot call the function on return command")
	}

	arg0 := strings.Split(p.GetCurrentLine(), " ")[0]
	if commandType == C_ARITHMETIC {
		return arg0
	}
	arg1 := strings.Split(p.GetCurrentLine(), " ")[1]

	return arg1
}

func (p *Parser) GetArg2() int {
	commandType := p.CommandType()
	if commandType != C_PUSH &&
		commandType != C_POP &&
		commandType != C_FUNCTION &&
		commandType != C_CALL {
		panic("cannot call the function on the command")
	}
	arg2 := strings.Split(p.GetCurrentLine(), " ")[2]
	intArg2, err := strconv.Atoi(arg2)
	if err != nil {
		panic("command is invalid")
	}
	return intArg2
}

func (p *Parser) GetCurrentLine() string {
	return strings.TrimSpace(string(p.file[p.currentCommad]))
}
