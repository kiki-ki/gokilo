package editor

type Code = byte

const (
	CodeHT  Code = 9
	CodeLF  Code = 10
	CodeCR  Code = 13
	CodeSP  Code = 32
	CodeDEL Code = 127
	// ALLOW_UP    Code = 1000
	// ALLOW_DOWN  Code = 1001
	// ALLOW_RIGHT Code = 1002
	// ALLOW_LEFT  Code = 1003
)
