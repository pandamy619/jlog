package jlog

import "os"

func (l *Jlogs) Info(message string) {
	l.typeConsoleLog(infoPrefix, message)
	l.outFile()
}

func (l *Jlogs) Warning(message string) {
	l.typeConsoleLog(warningPrefix, message)
	l.outFile()
}

func (l *Jlogs) Error(message string) {
	l.typeConsoleLog(errorPrefix, message)
	l.outFile()
	os.Exit(1)
}
