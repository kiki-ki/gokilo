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

func (termios *Termios) EnableRawMode() error {
	termios.termios.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON
	termios.termios.Oflag &^= unix.OPOST
	termios.termios.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN
	termios.termios.Cflag &^= unix.CSIZE | unix.PARENB
	termios.termios.Cflag |= unix.CS8
	termios.termios.Cc[unix.VMIN] = 1
	termios.termios.Cc[unix.VTIME] = 0

	if err := unix.IoctlSetTermios(int(os.Stdin.Fd()), unix.TIOCSETA, termios.termios); err != nil {
		return err
	}

	return nil
}

func (termios *Termios) DisableRawMode() error {
	if err := unix.IoctlSetTermios(int(os.Stdin.Fd()), unix.TIOCSETA, termios.termios); err != nil {
		return err
	}

	return nil
}
