package translator

import (
	"github.com/Farhan-slurrp/nand2tetris/project8/codewriter"
)

type Translator struct {
	writers []codewriter.ICodeWriter
}

type ITranslator interface {
	Translate(codewriter.ICodeWriter)
	TranslateAll()
}

func NewTranslator(filenames []string) ITranslator {
	var writers []codewriter.ICodeWriter
	for _, filename := range filenames {
		newWriter := codewriter.NewCodeWriter(filename)
		writers = append(writers, newWriter)
	}

	return &Translator{
		writers: writers,
	}
}

func (t *Translator) Translate(writer codewriter.ICodeWriter) {
	writer.Write()
	defer writer.Close()
}

func (t *Translator) TranslateAll() {
	for _, writer := range t.writers {
		t.Translate(writer)
	}
}
