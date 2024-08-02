package parser

import "github.com/Farhan-slurrp/nand2tetris/project6/symboltable"

type Parser struct {
	file               [][]byte
	currentInstruction int
	currentLine        string
	table              symboltable.ISymbolTable
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
