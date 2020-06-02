package logging

import (
	"fmt"
	"io/ioutil"
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

func (l *Logs) Info(message string){
	l.typeConsoleLog(infoPrefix, message)
}

func (l *Logs) Warning(message string) {
	l.typeConsoleLog(warningPrefix, message)
}

func (l *Logs) Error(message string) {
	l.typeConsoleLog(errorPrefix, message)
	// os.Exit(1)
}