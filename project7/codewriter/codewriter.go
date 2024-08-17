package codewriter

import (
	"fmt"
	"os"

	"github.com/Farhan-slurrp/nand2tetris/compiler/parser"
)

type CodeWriter struct {
	parser     parser.IParser
	outputFile *os.File
}

type ICodeWriter interface {
	writeArithmetic(command string)
	writePushPop(command parser.CmdType, segment string, index int)
	close()
}

func NewCodeWriter(parser parser.IParser, outputFile *os.File) ICodeWriter {
	// clear the output file
	outputFile.Truncate(0)
	outputFile.Seek(0, 0)
	return &CodeWriter{
		parser:     parser,
		outputFile: outputFile,
	}
}

func (cw *CodeWriter) writeArithmetic(command string) {
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

func (cw *CodeWriter) writePushPop(command parser.CmdType, segment string, index int) {
	switch command {
	case "push":
		cw.writePush()
	case "pop":
		cw.writePop()
	}
}

func (cw *CodeWriter) close() {
	cw.outputFile.Close()
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
	cw.writeFile(lineToWrite, "")
	if jumpType != "" {
		cw.jump(jumpType)
	}
	cw.pushStack(nil)
}

func (cw *CodeWriter) jump(jumpType string) {
	cw.writeFile("@TRUE_JUMP", "@")
	cw.writeFile("D; "+jumpType+"\nD=0", "")
	cw.writeFile("@FALSE_JUMP", "@")
	cw.writeFile("0;JMP", "")
	cw.writeFile("(TRUE_JUMP", "(")
	cw.writeFile("D=-1", "")
	cw.writeFile("(FALSE_JUMP", "(")
}

func (cw *CodeWriter) loadMemory() {
	// TO DO
}

func (cw *CodeWriter) loadStatic(pop bool) {
	arg2 := cw.parser.GetArg2()
	cw.writeFile(fmt.Sprintf("@%d", arg2), "")
	if pop {
		cw.writeFile("M=D", "")
	} else {
		cw.writeFile("D=M", "")
	}
}

func (cw *CodeWriter) popStack(saveToD bool) {
	cw.writeFile("@SP\nM=M-1\nA=M", "")
	if saveToD {
		cw.writeFile("D=M\n", "")
	}
}

func (cw *CodeWriter) pushStack(num *int) {
	if num != nil {
		cw.writeFile(fmt.Sprintf("@%d\nD=A", num), "")
	}
	cw.writeFile("@SP\nA=M\nM=D\n@SP\nM=M+1", "")
}

func (cw *CodeWriter) writeFile(str string, label string) {
	if label == "@" {
		cw.outputFile.WriteString(fmt.Sprintf("%s\n", str))
	} else {
		cw.outputFile.WriteString(fmt.Sprintf("%s)\n", str))
	}
}

func (cw *CodeWriter) writePop() {
	// TO DO
}

func (cw *CodeWriter) writePush() {
	arg1 := cw.parser.GetArg1()
	arg2 := cw.parser.GetArg2()

	switch arg1 {
	case "constant":
		cw.pushStack(&arg2)
	case "static":
		cw.loadStatic(false)
		cw.pushStack(nil)
	default:
		cw.loadMemory()
		cw.pushStack(nil)
	}
}
