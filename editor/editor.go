package editor

import (
	"bufio"
	"fmt"
	"os"
	"time"

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
	Status     *editorStatus
}

type editorStatus struct {
	Msg  string
	Time time.Time
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
		Status:     &editorStatus{},
	}

	e.setStatus("HELP: Ctrl-S = save | Ctrl-Q = quit | Ctrl-F = find")

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
	fp, err := os.Open(e.filePath)
	if err != nil {
		return err
	}
	defer fp.Close()

	sc := bufio.NewReader(fp)

	for l, err := sc.ReadBytes(CodeLF); err == nil; l, err = sc.ReadBytes(CodeLF) {
		for c := l[len(l)-1]; len(l) > 0 && (c == CodeLF || c == CodeCR); {
			// ここから
			l = l[:len(l)-1]
			if len(l) > 0 {
				c = l[len(l)-1]
			}
			e.insertRow(l)
		}
		fmt.Println(string(l))
	}

	return nil
}

func (e *Editor) insertRow(row []byte) {

}

func (e *Editor) setStatus(msgArgs ...interface{}) {
	e.Status.Msg = fmt.Sprintf(msgArgs[0].(string), msgArgs[:1]...)
	e.Status.Time = time.Now()
}

func getWindowSize() (*unix.Winsize, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, err
	}

	return ws, nil
}
