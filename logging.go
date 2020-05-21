package logging

import (
	"io"
)

type Logger struct {
	out io.Writer
}

func New(out io.Writer) *Logger {
	return &Logger{out: out}
}