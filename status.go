package jlog

import (
	"fmt"
	"os"
	"time"
)

func (l *Jlogs) makePath() string {
	return fmt.Sprintf("%s/%s", l.location, l.Pkg)
}

func (l *Jlogs) makeFilename() string {
	return fmt.Sprintf("%s/%s/%s.json", l.location, l.Pkg, time.Now().Format("2006-01-02T15:04:05"))
}

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
