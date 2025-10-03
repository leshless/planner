package stdio

import (
	"io"
	"os"
)

type StdIO interface {
	StdIn() io.Reader
	StdOut() io.Writer
	StdErr() io.Writer
}

type stdIO struct{}

var _ StdIO = (*stdIO)(nil)

func NewStdIO() StdIO {
	return &stdIO{}
}

func (s *stdIO) StdIn() io.Reader {
	return os.Stdin
}

func (s *stdIO) StdOut() io.Writer {
	return os.Stdout
}

func (s *stdIO) StdErr() io.Writer {
	return os.Stderr
}
