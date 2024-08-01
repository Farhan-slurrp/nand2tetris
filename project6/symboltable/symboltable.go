package symboltable

func NewSymbolTable() ISymbolTable {
	return &SymbolTable{
		table: make(map[string]int),
	}
}

func (st *SymbolTable) addEntry(symbol string, address int) {
	st.table[symbol] = address
}

func (st *SymbolTable) contains(key string) bool {
	_, ok := st.table[key]
	return ok
}

func (st *SymbolTable) getAddress(key string) int {
	val, ok := st.table[key]
	if !ok {
		return -1
	}
	return val
}
