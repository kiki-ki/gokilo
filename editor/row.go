package editor

type Row struct {
	chars []rune
}

func NewRow() *Row {
	return &Row{}
}
