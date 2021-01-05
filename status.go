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

func createStatusMessage(prefix string, message string) string {
	return fmt.Sprintf("%s %s %s", time.Now().Format(timeFormat), prefix, message)
}

// Info calls l.typeConsoleLog to print to the logger and print to the console
func (l *Jlogs) Info(message string) {
	l.consoleLog(infoPrefix, message)
	outFile(l.makePath(), l.makeFilename(), []byte(createStatusMessage("INFO", message)))
}

// Warning calls l.typeConsoleLog to print to the logger and print to the console
func (l *Jlogs) Warning(message string) {
	l.consoleLog(warningPrefix, message)
	outFile(l.makePath(), l.makeFilename(), []byte(createStatusMessage("WARN", message)))
}

// Error calls l.typeConsoleLog to print to the logger and print to the console.
// Followed by a call to os.Exit(1)
func (l *Jlogs) Error(message string) {
	l.consoleLog(errorPrefix, message)
	outFile(l.makePath(), l.makeFilename(), []byte(createStatusMessage("ERROR", message)))

	os.Exit(1)
}
