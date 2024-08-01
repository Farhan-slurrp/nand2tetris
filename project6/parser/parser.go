package parser

import (
	"fmt"
	"strings"
)

func NewParser(file [][]byte) IParser {
	return &Parser{
		file:               file,
		currentInstruction: -1,
		currentLine:        "",
	}
}

func (p *Parser) HasMoreLines() bool {
	return p.currentInstruction+1 < len(p.file)
}

func (p *Parser) Advance() {
	if p.HasMoreLines() {
		p.currentInstruction += 1
		p.currentLine = string(p.file[p.currentInstruction])
		if strings.HasPrefix(p.currentLine, "//") || len(p.currentLine) == 0 {
			return
		}
		instructionType := p.InstructionType()
		fmt.Println(instructionType, p.currentLine)
	}
}

func (p *Parser) InstructionType() string {
	if p.currentLine == "" {
		return ""
	} else if strings.HasPrefix(p.currentLine, "@") {
		return "A_INSTRUCTION"
	} else if strings.HasPrefix(p.currentLine, "(") && strings.HasSuffix(p.currentLine, ")") {
		return "L_INSTRUCTION"
	}
	return "C_INSTRUCTION"
}

func (p *Parser) Symbol() string {
	return ""
}

func (p *Parser) Dest() string {
	return ""
}

func (p *Parser) Comp() string {
	return ""
}

func (p *Parser) Jump() string {
	return ""
}
