package editor

import (
	"os"

	"golang.org/x/sys/unix"
)

type Termios struct {
	termios *unix.Termios
}

func newTermios() (*Termios, error) {
	termios, err := unix.IoctlGetTermios(int(os.Stdin.Fd()), unix.TIOCGETA)
	if err != nil {
		return nil, err
	}

	return &Termios{termios: termios}, nil
}

func (t *Termios) EnableRawMode() error {
	t.termios.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON
	t.termios.Oflag &^= unix.OPOST
	t.termios.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN
	t.termios.Cflag &^= unix.CSIZE | unix.PARENB
	t.termios.Cflag |= unix.CS8
	t.termios.Cc[unix.VMIN] = 1
	t.termios.Cc[unix.VTIME] = 0

	if err := unix.IoctlSetTermios(int(os.Stdin.Fd()), unix.TIOCSETA, t.termios); err != nil {
		return err
	}

	return nil
}

func (t *Termios) DisableRawMode() error {
	if err := unix.IoctlSetTermios(int(os.Stdin.Fd()), unix.TIOCSETA, t.termios); err != nil {
		return err
	}

	return nil
}
