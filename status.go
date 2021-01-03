package jlog

import "os"

const (
	infoPrefix    = infoColor + "[INFO]" + resetColor
	warningPrefix = warningColor + "[WARN]" + resetColor
	errorPrefix   = errorColor + "[ERROR]" + resetColor
)

// Info output to console and write to file, message
func (l *Jlogs) Info(message string) {
	l.typeConsoleLog(infoPrefix, message)
	outFile(l.makePath(), l.makeFilename(), l.Report())
}

// Warning output to console and write to file, message
func (l *Jlogs) Warning(message string) {
	l.typeConsoleLog(warningPrefix, message)
	outFile(l.makePath(), l.makeFilename(), l.Report())
}

// Error output to console and write to file, message after exit program
func (l *Jlogs) Error(message string) {
	l.typeConsoleLog(errorPrefix, message)
	outFile(l.makePath(), l.makeFilename(), l.Report())

	os.Exit(1)
}
