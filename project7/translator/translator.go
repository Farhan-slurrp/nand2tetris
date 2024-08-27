package translator

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Farhan-slurrp/nand2tetris/compiler/codewriter"
	"github.com/Farhan-slurrp/nand2tetris/compiler/parser"
)

type Translator struct {
	writer codewriter.ICodeWriter
}

type ITranslator interface {
	Translate()
}

func NewTranslator(filename string) ITranslator {
	file, err := os.ReadFile(fmt.Sprintf("./vm/%s.vm", filename))
	if err != nil {
		fmt.Println("Can't read file:", fmt.Sprintf("./vm/%s.vm", filename))
		panic(err)
	}

	array := bytes.Split(file, []byte("\n"))
	parser := parser.NewParser(array)
	writer := codewriter.NewCodeWriter(parser, filename)

	return &Translator{
		writer: writer,
	}
}

func (t *Translator) Translate() {
	t.writer.Write()
	defer t.writer.Close()
}
