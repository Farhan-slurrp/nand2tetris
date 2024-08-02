package symboltable

type SymbolTable struct {
	table map[string]int
}

type ISymbolTable interface {
	AddEntry(symbol string, address int)
	Contains(key string) bool
	GetAddress(key string) int
	MapKey(value int) (string, bool)
}
