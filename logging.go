package logging

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const (
	infoPrefix = "\u001B[34;1m[INFO]\u001B[0m"
	warningPrefix = "\u001B[33m[WARNING]\u001b[0m"
	errorPrefix = "\u001b[31m[ERROR]\u001b[0m"
)

func (l *Logs) outFile() {
	_ = ioutil.WriteFile(
		fmt.Sprintf("%s/%s.json", l.location, l.Pkg),
		l.Report(),
		0777,
	)
}
func (l *Logs) outputConsole(prefix string, row string) {
	datetime := time.Now().Format("2006-01-02T15:04:05")
	fmt.Printf("%s \u001b[32mPACKAGE\u001B[0m: %s %s %s\n",datetime, l.Pkg, prefix, row)
	l.outFile()
}

func (l *Logs) Info(message string) *Logs{
	l.outputConsole(infoPrefix, message)
	return &Logs{location: l.location, Pkg: l.Pkg, Funs: l.Funs,  name: l.name, log: l.log}
}

func (l *Logs) Warning(message string) *Logs{
	l.outputConsole(warningPrefix, message)
	return &Logs{location: l.location, Pkg: l.Pkg, Funs: l.Funs,  name: l.name, log: l.log}
}

func (l *Logs) Error(message string) {
	l.outputConsole(errorPrefix, message)
	os.Exit(1)
}