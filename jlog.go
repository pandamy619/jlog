package jlog

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const (
	infoPrefix = "\u001B[34;1m[INFO]\u001B[0m"
	warningPrefix = "\u001B[33m[WARN]\u001b[0m"
	errorPrefix = "\u001b[31m[ERROR]\u001b[0m"
)


func (l *Logs) checkDir(path string) {
	_, err := os.Stat(path)
	l.createDirAll(path, err)
}

func (l *Logs) createDirAll(path string, err error) {
	if os.IsNotExist(err) {
		_ = os.MkdirAll(path, 0777)
	}
}

func (l *Logs) outFile() {
	path := fmt.Sprintf("%s/%s", l.location, l.Pkg)
	l.checkDir(path)
	// TODO временное решение
	filename := time.Now().Format("2006-01-02T15:04:05")
	_ = ioutil.WriteFile(
		fmt.Sprintf("%s/%s/%s.json", l.location, l.Pkg, filename),
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
	os.Exit(1)
}