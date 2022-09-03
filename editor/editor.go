package editor

import (
	"fmt"
	"os"
)

type Editor struct {
	filePath string
	keyChan  chan rune
	cx       int
	cy       int
	rows     []*Row
	terminal *Terminal
}

func NewEditor(filePath string) (*Editor, error) {
	e := initEditor(filePath)

	fs, err := os.Stat(filePath)
	if err != nil {
		// loadEditor(filepath)
	}

	if fs.IsDir() {
		return nil, fmt.Errorf("%s is a directory", fs.Name())
	}

	return e, nil
}

func initEditor(filePath string) *Editor {
	e := &Editor{
		filePath: filePath,
		keyChan:  make(chan rune),
		cx:       0,
		cy:       0,
		rows:     make([]*Row, 0),
		terminal: newTerminal(),
	}

	return e
}
