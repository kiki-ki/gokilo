package editor

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

type Editor struct {
	filePath   string
	cx         int
	cy         int
	rows       []*row
	screenRows int
	screenCols int
	Termios    *Termios
}

func NewEditor(filePath string) (*Editor, error) {
	termios, err := newTermios()
	if err != nil {
		return nil, err
	}

	if err := termios.EnableRawMode(); err != nil {
		return nil, err
	}

	ws, err := getWindowSize()
	if err != nil {
		return nil, err
	}

	e := &Editor{
		filePath:   filePath,
		cx:         0,
		cy:         0,
		rows:       make([]*row, 0),
		screenRows: int(ws.Row) - 2,
		screenCols: int(ws.Col),
		Termios:    termios,
	}

	fs, err := os.Stat(e.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return e, nil
		}
		return nil, err
	}

	if fs.IsDir() {
		return nil, fmt.Errorf("%s is a directory", fs.Name())
	}

	if err := e.loadFile(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Editor) loadFile() error {
	// bytes, err := ioutil.ReadFile(e.filePath)
	// if err != nil {
	// 	return err
	// }

	// for _, b := range bytes {
	// 	e.rows = append(e.rows)
	// }

	return nil
}

func getWindowSize() (*unix.Winsize, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, err
	}

	return ws, nil
}
