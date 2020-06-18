package jlog

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Jlogs struct {
	location string
	Pkg string `json:"package"`
	Funs map[string]interface{} `json:"functions"`
	log []interface{}
	name string
	fields Fields
	consoleLog string
}

func Init(pkg string, location string, consoleLog string) *Jlogs {
	m := make(map[string]interface{})
	return &Jlogs{
		location: location,
		Pkg: pkg,
		Funs: m,
		consoleLog: consoleLog,
	}
}

func (l *Jlogs) Main(name string) *Jlogs {
	l.name = name
	return &Jlogs{
		location: l.location,
		Pkg: l.Pkg,
		Funs: l.Funs,
		name: l.name,
		log: l.log,
		fields: l.fields,
		consoleLog: l.consoleLog,
	}
}

func (l *Jlogs) Sub(status string, message string, time string) *Jlogs {
	f := Fields{
		"status": status,
		"message": message,
		"time": time,
	}
	l.fields = f
	return l.updateArr(f)
}

func (l *Jlogs) SubWithFields(status string, message string, time string, fields Fields) *Jlogs {
	f := Fields{
		"status": status,
		"message": message,
		"time": time,
		"meta": fields,
	}
	l.fields = f
	return l.updateArr(f)
}

func (l *Jlogs) updateArr(fields Fields) *Jlogs {
	l.log = append(l.log, fields)
	l.Funs[l.name] = l.log
	return &Jlogs{
		location: l.location,
		Pkg: l.Pkg,
		Funs: l.Funs,
		name: l.name,
		log: l.log,
		fields: l.fields,
		consoleLog: l.consoleLog,
	}
}

func (l *Jlogs) Report() []byte{
	obj := structJson(l)
	return obj
}

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