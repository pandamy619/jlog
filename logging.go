package logging

import (
	"io"
)

type Logger struct {
	// Time time.Time
	// FormatTime string
	out io.Writer
	fields string
}

// NewLog creates a new Logger
func NewLog(out io.Writer) *Logger {
	return &Logger{out: out}
}

func (l *Logger) checkLineBreak(buf []byte) []byte{
	if len(buf) == 0 || buf[len(buf) - 1] != '\n' {
		buf = append(buf, '\n')
	}
	return buf
}

func (l *Logger) convertToByte(row string) []byte{
	buf := []byte(row)
	return l.checkLineBreak(buf)
}

func (l *Logger) Output(row string) error{
	buf := l.convertToByte(row)
	_, err := l.out.Write(buf)
	return err
}