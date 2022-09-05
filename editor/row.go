package editor

type row struct {
	chars []rune
}

func newRow() *row {
	return &row{}
}
