package main

import (
	"os"

	"github.com/Farhan-slurrp/nand2tetris/compiler/translator"
)

func main() {
	if len(os.Args) < 2 {
		panic("missing parameter, provide filename")
	}

	translator := translator.NewTranslator(os.Args[1])
	translator.Translate()
}
