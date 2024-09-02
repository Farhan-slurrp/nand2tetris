package translator

import (
	"github.com/Farhan-slurrp/nand2tetris/project8/codewriter"
)

type Translator struct {
	writer codewriter.ICodeWriter
}

type ITranslator interface {
	Translate()
}

func NewTranslator(filename string) ITranslator {
	writer := codewriter.NewCodeWriter(filename)

	return &Translator{
		writer: writer,
	}
}

func (t *Translator) Translate() {
	t.writer.Write()
	defer t.writer.Close()
}
