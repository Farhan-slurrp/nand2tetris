package parser

import (
	"os"

	"github.com/Farhan-slurrp/nand2tetris/project6/symboltable"
)

type Parser struct {
	file               [][]byte
	currentInstruction int
	currentLine        string
	outputFile         *os.File
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
	FirstScan()
}
