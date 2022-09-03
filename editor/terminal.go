package editor

import "golang.org/x/sys/unix"

type Terminal struct {
	termios *unix.Termios
	width   int
	height  int
}

func newTerminal() *Terminal {
	return &Terminal{}
}
