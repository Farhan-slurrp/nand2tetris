package code

func NewCode(dest string, comp string, jump string) ICode {
	return &Code{
		destInstruction: dest,
		compInstruction: comp,
		jumpInstruction: jump,
	}
}

func (c *Code) Dest() string {
	instructionCode, ok := DestLookup[c.destInstruction]
	if !ok {
		return ""
	}
	return instructionCode
}

func (c *Code) Comp() (val string, a string) {
	instructionCode, ok := CompLookup[c.compInstruction]
	if !ok {
		return "", ""
	}
	return instructionCode.val, instructionCode.a
}

func (c *Code) Jump() string {
	instructionCode, ok := JumpLookup[c.jumpInstruction]
	if !ok {
		return ""
	}
	return instructionCode
}
