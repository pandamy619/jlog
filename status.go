package jlog

import "os"

func (l *Jlogs) Info(message string) {
	l.typeConsoleLog(infoPrefix, message)
	outFile(l.makePath(), l.makeFilename(), l.Report())
}

func (l *Jlogs) Warning(message string) {
	l.typeConsoleLog(warningPrefix, message)
	outFile(l.makePath(), l.makeFilename(), l.Report())
}

func (l *Jlogs) Error(message string) {
	l.typeConsoleLog(errorPrefix, message)
	outFile(l.makePath(), l.makeFilename(), l.Report())

	os.Exit(1)
}
