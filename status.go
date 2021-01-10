package jlog

import (
	"fmt"
	"os"
	"time"
)

const (
	infoPrefix    = infoColor + "[INFO]" + resetColor
	warningPrefix = warningColor + "[WARN]" + resetColor
	errorPrefix   = errorColor + "[ERROR]" + resetColor
)

func (l *Jlogs) output(status string, prefix string, message ...interface{}) {
	m := fmt.Sprint(message...)
	l.consoleLog(prefix, m)
	// TODO добавить наименование пакета и функции
	b := []byte(fmt.Sprintf("%s %s %s\n", time.Now().Format(timeFormat), status, m))
	outFile(l.makePath(), l.makeFilename(), b)
}

// Info calls l.typeConsoleLog to print to the logger and print to the console
func (l *Jlogs) Info(message ...interface{}) {
	l.output("INFO", infoPrefix, message...)
}

// Warning calls l.typeConsoleLog to print to the logger and print to the console
func (l *Jlogs) Warning(message ...interface{}) {
	l.output("WARN", warningPrefix, message...)
}

// Error calls l.typeConsoleLog to print to the logger and print to the console.
// Followed by a call to os.Exit(1)
func (l *Jlogs) Error(message ...interface{}) {
	l.output("ERROR", errorPrefix, message...)
	os.Exit(1)
}
