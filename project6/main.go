package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Farhan-slurrp/nand2tetris/project6/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide filename!")
		return
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	array := bytes.Split(file, []byte("\n"))

	my_parser := parser.NewParser(array)
	for {
		if my_parser.HasMoreLines() == false {
			break
		}
		my_parser.Advance()
	}
}
