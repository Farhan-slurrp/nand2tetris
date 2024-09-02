package codewriter

import (
	"fmt"
	"os"
	"strings"

	"github.com/Farhan-slurrp/nand2tetris/project8/parser"
)

type CodeWriter struct {
	parser     parser.IParser
	outputFile *os.File
}

type ICodeWriter interface {
	Write()
	Close()
}

var SymbolHash = map[string]string{
	"local":    "LOCAL",
	"argument": "ARG",
	"this":     "THIS",
	"that":     "THAT",
	"pointer":  "THIS",
	"temp":     "5",
}

func NewCodeWriter(filename string) ICodeWriter {
	parser := parser.NewParser(filename)

	out, err := os.OpenFile(fmt.Sprintf("./asm/%s.asm", filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Can't read file:", fmt.Sprintf("./asm/%s.asm", filename))
		panic(err)
	}
	// clear the output file
	out.Truncate(0)
	out.Seek(0, 0)
	return &CodeWriter{
		parser:     parser,
		outputFile: out,
	}
}

func (cw *CodeWriter) Close() {
	cw.outputFile.Close()
}

func (cw *CodeWriter) Write() {
	for cw.parser.HasMoreLines() {
		currentLine := cw.parser.GetCurrentLine()
		if !strings.HasPrefix((currentLine), "//") &&
			len(currentLine) != 0 {
			cw.translate()
		}
		cw.parser.Advance()
	}
}

func (cw *CodeWriter) translate() {
	cmd := cw.parser.CommandType()
	switch cmd {
	case parser.C_ARITHMETIC:
		cw.writeArithmetic()
	case parser.C_POP:
		cw.writePop()
	case parser.C_PUSH:
		cw.writePush()
	}
}

func (cw *CodeWriter) arithmetic(symbol string, jumpType string, unary bool) {
	cw.popStack(true)
	if !unary {
		cw.popStack(false)
	}
	lineToWrite := "D="
	if !unary {
		lineToWrite += "M"
	}
	lineToWrite += symbol + "D"
	cw.writeFile(lineToWrite)
	if jumpType != "" {
		cw.jump(jumpType)
	}
	cw.pushStack(-1)
}

func (cw *CodeWriter) jump(jumpType string) {
	cw.writeFile("@TRUE_JUMP")
	cw.writeFile("D; " + jumpType + "\nD=0")
	cw.writeFile("@FALSE_JUMP")
	cw.writeFile("0;JMP")
	cw.writeFile("(TRUE_JUMP")
	cw.writeFile("D=-1")
	cw.writeFile("(FALSE_JUMP")
}

func (cw *CodeWriter) loadMemory(saveFromR13 bool) {
	arg1 := cw.parser.GetArg1()
	arg2 := cw.parser.GetArg2()
	cw.writeFile(fmt.Sprintf("@%d", arg2))
	cw.writeFile("D=A")
	cw.writeFile(SymbolHash[arg1])

	if arg1 == "temp" || arg1 == "pointer" {
		cw.writeFile("AD=A+D")
	} else {
		cw.writeFile("AD=M+D")
	}

	if saveFromR13 {
		cw.writeFile("@14\nM=D\n@13\nD=M\n@14\nA=M\nM=D")
	} else {
		cw.writeFile("D=M")
	}
}

func (cw *CodeWriter) loadStatic(pop bool) {
	arg2 := cw.parser.GetArg2()
	cw.writeFile(fmt.Sprintf("@%d", arg2))
	if pop {
		cw.writeFile("M=D")
	} else {
		cw.writeFile("D=M")
	}
}

func (cw *CodeWriter) popStack(saveToD bool) {
	cw.writeFile("@SP\nM=M-1\nA=M")
	if saveToD {
		cw.writeFile("D=M")
	}
}

func (cw *CodeWriter) pushStack(num int) {
	if num >= 0 {
		cw.writeFile(fmt.Sprintf("@%d\nD=A", num))
	}
	cw.writeFile("@SP\nA=M\nM=D\n@SP\nM=M+1")
}

func (cw *CodeWriter) writeFile(str string) {
	cw.outputFile.WriteString(fmt.Sprintf("%s\n", str))
}

func (cw *CodeWriter) writeArithmetic() {
	arg0 := cw.parser.GetArg1()
	switch arg0 {
	case "add":
		cw.arithmetic("+", "", false)
	case "sub":
		cw.arithmetic("-", "", false)
	case "eq":
		cw.arithmetic("-", "JEQ", false)
	case "lt":
		cw.arithmetic("-", "JLT", false)
	case "gt":
		cw.arithmetic("-", "JGT", false)
	case "and":
		cw.arithmetic("&", "", false)
	case "or":
		cw.arithmetic("|", "", false)
	case "neg":
		cw.arithmetic("-", "", true)
	case "not":
		cw.arithmetic("!", "", true)
	}
}

func (cw *CodeWriter) writePop() {
	cw.popStack(true)
	arg1 := cw.parser.GetArg1()
	if arg1 == "parser" {
		cw.loadStatic(true)
	} else {
		cw.writeFile("@13\nM=D")
		cw.loadMemory(true)
	}

}

func (cw *CodeWriter) writePush() {
	arg1 := cw.parser.GetArg1()
	arg2 := cw.parser.GetArg2()

	switch arg1 {
	case "constant":
		cw.pushStack(arg2)
	case "static":
		cw.loadStatic(false)
		cw.pushStack(-1)
	default:
		cw.loadMemory(false)
		cw.pushStack(-1)
	}
}

//lint:ignore U1000 unused function
func (cw *CodeWriter) writeLabel(label string) {}

//lint:ignore U1000 unused function
func (cw *CodeWriter) writeGoto(label string) {}

//lint:ignore U1000 unused function
func (cw *CodeWriter) writeIf(label string) {}

//lint:ignore U1000 unused function
func (cw *CodeWriter) writeFunction(label string) {}

//lint:ignore U1000 unused function
func (cw *CodeWriter) writeReturn(label string) {}
