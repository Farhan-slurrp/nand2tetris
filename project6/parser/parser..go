package parser

func NewParser(filename string) IParser {
	return &Parser{
		filename:            filename,
		current_instruction: nil,
	}
}

func (p *Parser) hasMoreLines() bool {
	return false
}

func (p *Parser) advance() {

}

func (p *Parser) instructionType() string {
	return ""
}

func (p *Parser) symbol() string {
	return ""
}

func (p *Parser) dest() string {
	return ""
}

func (p *Parser) comp() string {
	return ""
}

func (p *Parser) jump() string {
	return ""
}
