package code

type CompReturn struct {
	val string
	a   string
}

var DestLookup = map[string]string{
	"":    "000",
	"M":   "001",
	"D":   "010",
	"DM":  "011",
	"MD":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"ADM": "111",
}

var CompLookup = map[string]CompReturn{
	"0": {
		val: "101010",
		a:   "0",
	},
	"1": {
		val: "111111",
		a:   "0",
	},
	"-1": {
		val: "111010",
		a:   "0",
	},
	"D": {
		val: "001100",
		a:   "0",
	},
	"A": {
		val: "110000",
		a:   "0",
	},
	"M": {
		val: "110000",
		a:   "1",
	},
	"!D": {
		val: "001101",
		a:   "0",
	},
	"!A": {
		val: "110001",
		a:   "0",
	},
	"!M": {
		val: "110001",
		a:   "1",
	},
	"D+1": {
		val: "011111",
		a:   "0",
	},
	"A+1": {
		val: "110111",
		a:   "0",
	},
	"M+1": {
		val: "110111",
		a:   "1",
	},
	"D-1": {
		val: "001110",
		a:   "0",
	},
	"A-1": {
		val: "110010",
		a:   "0",
	},
	"M-1": {
		val: "110010",
		a:   "1",
	},
	"D+A": {
		val: "000010",
		a:   "0",
	},
	"D+M": {
		val: "000010",
		a:   "1",
	},
	"D-A": {
		val: "010011",
		a:   "0",
	},
	"D-M": {
		val: "010011",
		a:   "1",
	},
	"A-D": {
		val: "000111",
		a:   "0",
	},
	"M-D": {
		val: "000111",
		a:   "1",
	},
	"D&A": {
		val: "000000",
		a:   "0",
	},
	"D&M": {
		val: "000000",
		a:   "1",
	},
	"D|A": {
		val: "010101",
		a:   "0",
	},
	"D|M": {
		val: "010101",
		a:   "1",
	},
}

var JumpLookup = map[string]string{
	"":    "000",
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}
