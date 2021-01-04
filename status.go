package jlog

import "os"

const (
	infoPrefix    = infoColor + "[INFO]" + resetColor
	warningPrefix = warningColor + "[WARN]" + resetColor
	errorPrefix   = errorColor + "[ERROR]" + resetColor
)

// Info calls l.typeConsoleLog to print to the logger and print to the console
func (l *Jlogs) Info(message string) {
	l.consoleLog(infoPrefix, message)
	outFile(l.makePath(), l.makeFilename(), l.Report())
}

// Warning calls l.typeConsoleLog to print to the logger and print to the console
func (l *Jlogs) Warning(message string) {
	l.consoleLog(warningPrefix, message)
	outFile(l.makePath(), l.makeFilename(), l.Report())
}

// Error calls l.typeConsoleLog to print to the logger and print to the console.
// Followed by a call to os.Exit(1)
func (l *Jlogs) Error(message string) {
	l.consoleLog(errorPrefix, message)
	outFile(l.makePath(), l.makeFilename(), l.Report())

	os.Exit(1)
}
