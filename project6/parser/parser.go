package parser

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Farhan-slurrp/nand2tetris/project6/code"
	"github.com/Farhan-slurrp/nand2tetris/project6/symboltable"
	"github.com/Farhan-slurrp/nand2tetris/project6/utils"
)

func NewParser(file [][]byte, table symboltable.ISymbolTable, outputFile *os.File) IParser {
	return &Parser{
		file:               file,
		currentInstruction: -1,
		currentLine:        "",
		table:              table,
		outputFile:         outputFile,
	}
}

func (p *Parser) HasMoreLines() bool {
	return p.currentInstruction+1 < len(p.file)
}

func (p *Parser) Advance() {
	if p.HasMoreLines() {
		p.currentInstruction += 1
		p.currentLine = strings.TrimSpace(string(p.file[p.currentInstruction]))
		if strings.HasPrefix(p.currentLine, "//") ||
			len(p.currentLine) == 0 ||
			len(p.currentLine) == 1 {
			return
		}
		instructionType := p.InstructionType()
		fmt.Println(instructionType, p.currentLine)
		switch it := instructionType; it {
		case "L_INSTRUCTION":
			symbol := p.Symbol()
			binaryCode := fmt.Sprintf("%016b", p.table.GetAddress(symbol))
			p.outputFile.Write([]byte(binaryCode))
			p.outputFile.WriteString("\n")
		case "A_INSTRUCTION":
			symbol := p.Symbol()
			if !p.table.Contains(symbol) {
				address := 16
				numberSymbol, err := strconv.Atoi(symbol)
				if err != nil {
					for {
						_, ok := p.table.MapKey(address)
						if !ok {
							break
						}

						address += 1
					}
				} else {
					address = numberSymbol
				}
				p.table.AddEntry(symbol, address)
			}
			binaryCode := fmt.Sprintf("%016b", p.table.GetAddress(symbol))
			p.outputFile.Write([]byte(binaryCode))
			p.outputFile.WriteString("\n")
		case "C_INSTRUCTION":
			code := code.NewCode(p.Dest(), p.Comp(), p.Jump())
			comp, a := code.Comp()
			dest := code.Dest()
			jump := code.Jump()
			binaryCode := fmt.Sprintf("111%s%s%s%s", a, comp, dest, jump)
			p.outputFile.Write([]byte(binaryCode))
			p.outputFile.WriteString("\n")
		default:
			return
		}

	}
}

func (p *Parser) InstructionType() string {
	utils.Assert(p.currentLine != "", fmt.Errorf("cannot call the function on empty string"))
	sanitizedCurrentLine := strings.TrimSpace(p.currentLine)
	if strings.HasPrefix(sanitizedCurrentLine, "@") {
		return "A_INSTRUCTION"
	} else if strings.HasPrefix(sanitizedCurrentLine, "(") || strings.HasSuffix(sanitizedCurrentLine, ")") {
		return "L_INSTRUCTION"
	}
	return "C_INSTRUCTION"
}

func (p *Parser) Symbol() string {
	instructionType := p.InstructionType()
	utils.Assert(
		instructionType == "A_INSTRUCTION" || instructionType == "L_INSTRUCTION",
		fmt.Errorf("cannot call the function on type besides A_INSTRUCTION or L_INSTRUCTION"),
	)
	sanitizedCurrentLine := strings.TrimSpace(p.currentLine)
	if instructionType == "A_INSTRUCTION" {
		return strings.TrimPrefix(sanitizedCurrentLine, "@")
	} else {
		trimmed := strings.TrimPrefix(sanitizedCurrentLine, "(")
		return strings.TrimSpace(strings.TrimSuffix(trimmed, ")"))
	}
}

func (p *Parser) Dest() string {
	instructionType := p.InstructionType()
	utils.Assert(instructionType == "C_INSTRUCTION", fmt.Errorf("function should be called on C_INSTRUCTION string"))

	splitted := strings.Split(p.currentLine, "=")
	if len(splitted) == 1 {
		return ""
	}
	return strings.TrimSpace(splitted[0])
}

func (p *Parser) Comp() string {
	instructionType := p.InstructionType()
	utils.Assert(instructionType == "C_INSTRUCTION", fmt.Errorf("function should be called on C_INSTRUCTION string"))

	splitted := strings.Split(p.currentLine, "=")
	if len(splitted) == 1 {
		splitted = strings.Split(splitted[0], ";")
	} else {
		splitted = strings.Split(splitted[1], ";")
	}
	return strings.TrimSpace(splitted[0])
}

func (p *Parser) Jump() string {
	instructionType := p.InstructionType()
	utils.Assert(instructionType == "C_INSTRUCTION", fmt.Errorf("function should be called on C_INSTRUCTION string"))

	splitted := strings.Split(p.currentLine, "=")
	if len(splitted) == 1 {
		splitted = strings.Split(splitted[0], ";")
	} else {
		splitted = strings.Split(splitted[1], ";")
	}
	if len(splitted) == 1 {
		return ""
	}
	return strings.TrimSpace(splitted[1])
}

// put L_INSTRUCTION(s) to table
func (p *Parser) FirstScan() {
	// clear the output file
	p.outputFile.Truncate(0)
	p.outputFile.Seek(0, 0)
	for idx := range p.file {
		p.currentInstruction = idx
		p.currentLine = strings.TrimSpace(string(p.file[p.currentInstruction]))
		if strings.HasPrefix(p.currentLine, "//") ||
			len(p.currentLine) == 0 ||
			len(p.currentLine) == 1 {
			continue
		}
		if p.InstructionType() == "L_INSTRUCTION" {
			symbol := p.Symbol()
			if !p.table.Contains(symbol) {
				address := p.currentInstruction + 1
				nextLineStr := string(p.file[address])
				for {
					if !strings.HasPrefix(nextLineStr, "//") && len(nextLineStr) > 1 {
						break
					}
					address += 1
					nextLineStr = string(p.file[address])
				}
				p.table.AddEntry(symbol, address+1)
			}
		}
	}
	p.currentInstruction = -1
	p.currentLine = ""
}
