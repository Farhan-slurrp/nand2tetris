package main

import (
	"os"
	"strings"

	"github.com/Farhan-slurrp/nand2tetris/project8/translator"
)

func main() {
	if len(os.Args) < 2 {
		panic("missing parameter, provide directory")
	}

	dir := os.Args[1]
	files, err := os.ReadDir("./test/" + dir)
	if err != nil {
		panic(err)
	}

	var filenames []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".vm") {
			filenames = append(filenames, strings.TrimSuffix(dir+"/"+file.Name(), ".vm"))
		}
	}
	translator := translator.NewTranslator(filenames)
	translator.TranslateAll()
}
