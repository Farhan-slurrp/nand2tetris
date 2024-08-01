package symboltable

type SymbolTable struct {
	table map[string]int
}

type ISymbolTable interface {
	addEntry(symbol string, address int)
	contains(key string) bool
	getAddress(key string) int
}
