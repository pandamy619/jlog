package jlog

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func (l *Jlogs) makePath() string {
	return fmt.Sprintf("%s/%s", l.location, l.Pkg)
}

func (l *Jlogs) makeFilename() string {
	return fmt.Sprintf("%s.json", time.Now().Format("2006-01-02T15:04:05"))
}

// Creates a directory named path
func createDirs(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(path, 0777)
	}
}

// Creates a directory named path and save file to this directory
func outFile(path string, filename string, makeData []byte) {
	createDirs(path)
	_ = ioutil.WriteFile(path + "/" + filename, makeData,0777)
}

