package logging

import (
	"fmt"
	"io"
)


/* palette
Black: \u001b[30m
Red: \u001b[31m
Green: \u001b[32m
Yellow: \u001b[33m
Blue: \u001b[34m
Magenta: \u001b[35m
Cyan: \u001b[36m
White: \u001b[37m
Reset: \u001b[0m
*/

const (
	infoPrefix = "\u001B[34;1m[INFO]\u001B[0m"
	warningPrefix = "\u001B[33m[WARNING]\u001b[0m"
	errorPrefix = "\u001b[31m[ERROR]\u001b[0m"
)


type Logger struct {
	time string
	out io.Writer
	fields string
}

// NewLog creates a new Logger
func NewLog(time string, out io.Writer) *Logger {
	return &Logger{time: time, out: out}
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

func (l *Logger) outputConsole(time string, prefix string, row string) {
	fmt.Printf("%s %s %s\n", time, prefix, row)
}

func (l *Logger) Info(message string) {
	row := fmt.Sprintf("%s %s", l.fields, message)
	l.outputConsole(l.time, infoPrefix, row)
}

func (l *Logger) Warning(message string) {
	row := fmt.Sprintf("%s %s", l.fields, message)
	l.outputConsole(l.time, warningPrefix, row)
}

func (l *Logger) Error(message string) {
	row := fmt.Sprintf("%s %s", l.fields, message)
	l.outputConsole(l.time, errorPrefix, row)
	// TODO после тестов убрать
	// os.Exit(1)
}