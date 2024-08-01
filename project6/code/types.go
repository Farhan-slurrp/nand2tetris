package code

type Code struct {
	instruction string
}

type ICode interface {
	dest() string
	comp() string
	jump() string
}
