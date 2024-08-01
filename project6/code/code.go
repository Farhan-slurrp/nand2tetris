package code

func NewCode(instruction string) ICode {
	return &Code{
		instruction: instruction,
	}
}

func (c *Code) dest() string {
	return ""
}

func (c *Code) comp() string {
	return ""
}

func (c *Code) jump() string {
	return ""
}
