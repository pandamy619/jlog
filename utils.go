package jlog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func structJson(obj interface{}) []byte{
	o, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return o
}

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

func (l *Jlogs) checkDir(path string) {
	_, err := os.Stat(path)
	l.createDirAll(path, err)
}

func (l *Jlogs) createDirAll(path string, err error) {
	if os.IsNotExist(err) {
		_ = os.MkdirAll(path, 0777)
	}
}
