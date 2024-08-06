package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

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

	// output file
	re := regexp.MustCompile("/[A-Za-z]+.asm")
	match := re.FindStringSubmatch(os.Args[1])
	utils.Assert(len(match) > 0, fmt.Errorf("invalid file type"))
	outputFile := strings.Split(match[0], ".")[0]
	outputFile = "./hack/" + outputFile + ".hack"
	out, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Can't read file:", outputFile)
		panic(err)
	}
	defer out.Close()

	array := bytes.Split(file, []byte("\n"))
	table := symboltable.NewSymbolTable()
	myParser := parser.NewParser(array, table, out)
	// run first scan
	myParser.FirstScan()
	for {
		if !myParser.HasMoreLines() {
			break
		}
		myParser.Advance()
	}
}
