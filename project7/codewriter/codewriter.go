package codewriter

import (
	"os"

	"github.com/Farhan-slurrp/nand2tetris/compiler/parser"
)

type CodeWriter struct {
	parser     *parser.IParser
	outputFile *os.File
}

type ICodeWriter interface {
	writeArithmetic()
	writePushPop()
	close()
}

func NewCodeWriter(parser *parser.IParser, outputFile *os.File) ICodeWriter {
	// clear the output file
	outputFile.Truncate(0)
	outputFile.Seek(0, 0)
	return &CodeWriter{
		parser:     parser,
		outputFile: outputFile,
	}
}

func (cw *CodeWriter) writeArithmetic() {
	// TO DO
	cw.arithmetic()
}

func (cw *CodeWriter) writePushPop() {
	// TO DO
}

func (cw *CodeWriter) close() {
	cw.outputFile.Close()
}

func (cw *CodeWriter) arithmetic() {
	// TO DO
	cw.popStack(true)
}

func (cw *CodeWriter) popStack(saveToD bool) {
	cw.outputFile.WriteString("@SP\n")
	cw.outputFile.WriteString("M=M-1\n")
	cw.outputFile.WriteString("A=M\n")
	if saveToD {
		cw.outputFile.WriteString("D=M\n")
	}
}
