package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Farhan-slurrp/nand2tetris/project6/parser"
	"github.com/Farhan-slurrp/nand2tetris/project6/symboltable"
	"github.com/Farhan-slurrp/nand2tetris/project6/utils"
)

func main() {
	utils.Assert(len(os.Args) >= 2, fmt.Errorf("missing parameter, provide filename"))

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	array := bytes.Split(file, []byte("\n"))
	table := symboltable.NewSymbolTable()
	myParser := parser.NewParser(array, table)
	for {
		if !myParser.HasMoreLines() {
			break
		}
		myParser.Advance()
	}
}
