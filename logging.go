package logging

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	infoPrefix = "\u001B[34;1m[INFO]\u001B[0m"
	warningPrefix = "\u001B[33m[WARNING]\u001b[0m"
	errorPrefix = "\u001b[31m[ERROR]\u001b[0m"
)

func (l *Logs) outFile() {
	_, err := os.Stat(l.location)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(l.location, 0777)
	}
	_ = ioutil.WriteFile(
		fmt.Sprintf("%s/%s.json", l.location, l.Pkg),
		l.Report(),
		0777,
	)
}

func (l *Logs) Info(message string){
	l.typeConsoleLog(infoPrefix, message)
	l.outFile()
}

func (l *Logs) Warning(message string) {
	l.typeConsoleLog(warningPrefix, message)
	l.outFile()
}

func (l *Logs) Error(message string) {
	l.typeConsoleLog(errorPrefix, message)
	l.outFile()
	// os.Exit(1)
}