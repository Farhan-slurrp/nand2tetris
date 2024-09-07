package codewriter

import (
	"fmt"
	"os"
	"strings"

	"github.com/Farhan-slurrp/nand2tetris/project8/parser"
)

type CodeWriter struct {
	parser        parser.IParser
	outputFile    *os.File
	functionCount int
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
		parser:        parser,
		outputFile:    out,
		functionCount: 1,
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
	case parser.C_LABEL:
		cw.writeLabel()
	case parser.C_GOTO:
		cw.writeGoto()
	case parser.C_IF:
		cw.writeIf()
	case parser.C_FUNCTION:
		cw.writeFunction()
	case parser.C_CALL:
		cw.writeCall(false)
	case parser.C_RETURN:
		cw.writeReturn()
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

func (cw *CodeWriter) functionInit(nArgs int) {
	cw.writeFile(fmt.Sprintf("@RETURN%d\nD=A", cw.functionCount))
	vars := []string{"THAT", "THIS", "ARG", "LCL"}
	for _, v := range vars {
		cw.writeFile(fmt.Sprintf("@%s\nD=M", v))
	}
	cw.writeFile(fmt.Sprintf("@%d\nD=A\n@SP\nD=M-D\n@ARG\nM=D\n@SP\nD=M\n@LCL\nM=D", nArgs+5))
	cw.functionCount += 1
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

func (cw *CodeWriter) writeLabel() {
	arg1 := cw.parser.GetArg1()
	cw.writeFile(fmt.Sprintf("(%s)", arg1))
}

func (cw *CodeWriter) writeGoto() {
	arg1 := cw.parser.GetArg1()
	cw.writeFile(fmt.Sprintf("@%s", arg1))
	cw.writeFile("0;JMP")
}

func (cw *CodeWriter) writeIf() {
	arg1 := cw.parser.GetArg1()
	cw.popStack(false)
	cw.writeFile(fmt.Sprintf("@%s", arg1))
	cw.writeFile("D;JNE")
}

func (cw *CodeWriter) writeFunction() {
	arg1 := cw.parser.GetArg1()
	arg2 := cw.parser.GetArg2()
	cw.writeFile(fmt.Sprintf("(%s)", arg1))

	for i := 0; i < arg2; i++ {
		cw.writeFile("@0\nD=A")
		cw.pushStack(-1)
	}
}

func (cw *CodeWriter) writeCall(init bool) {
	arg1 := cw.parser.GetArg1()
	arg2 := cw.parser.GetArg2()
	nArgs := arg2
	if init {
		nArgs = 0
	}
	cw.functionInit(nArgs)
	if init {
		cw.writeFile("Sys.init")
	} else {
		cw.writeFile(fmt.Sprintf("%s\n0;JMP", arg1))
	}
	cw.writeFile(fmt.Sprintf("(RETURN%d)", cw.functionCount-1))
}

func (cw *CodeWriter) writeReturn() {
	cw.writeFile("@5\nD=A\n@LCL\nA=M-D\nD=M\n@15\nM=D")
	cw.popStack(false)
	cw.writeFile("@ARG\nA=M\nM=D\nD=A+1\n@SP\nM=D")
	vars := []string{"THAT", "THIS", "ARG", "LCL"}
	for _, v := range vars {
		cw.writeFile(fmt.Sprintf("@LCL\nAM=M-1\nD=M\n@%s\nM=D", v))
	}
	cw.writeFile("@15\nA=M")
	cw.writeFile("0;JMP")
}
