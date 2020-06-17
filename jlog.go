package jlog

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)


// TODO в другой файл (example: workFile)
func (l *Jlogs) outFile() {
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

func (l *Jlogs) Info(message string){
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