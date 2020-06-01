package logging

import (
	"fmt"
	"io/ioutil"
)

const (
	infoPrefix = "\u001B[34;1m[INFO]\u001B[0m"
	warningPrefix = "\u001B[33m[WARNING]\u001b[0m"
	errorPrefix = "\u001b[31m[ERROR]\u001b[0m"
	infoColor = "\u001B[34;1m"
	warningColor = "\u001B[33m"
	errorColor = "\u001B[31m"
	blockColor = "\u001B[0m"
)

func (l *Logs) outFile() {
	_ = ioutil.WriteFile(
		fmt.Sprintf("%s/%s.json", l.location, l.Pkg),
		l.Report(),
		0777,
	)
}

func (l *Logs) Info(message string) *Logs{
	l.consoleLog(l.Pkg, l.name, infoPrefix, message)
	return &Logs{location: l.location, Pkg: l.Pkg, Funs: l.Funs,  name: l.name, log: l.log}
}

func (l *Logs) Warning(message string) *Logs{
	l.consoleLog(l.Pkg, l.name, warningPrefix, message)
	return &Logs{location: l.location, Pkg: l.Pkg, Funs: l.Funs,  name: l.name, log: l.log}
}

func (l *Logs) Error(message string) {
	l.consoleLog(l.Pkg, l.name, errorPrefix, message)
	// os.Exit(1)
}