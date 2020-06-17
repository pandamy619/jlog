package jlog

import "os"


func (l *Jlogs) checkDir(path string) {
	_, err := os.Stat(path)
	l.createDirAll(path, err)
}

func (l *Jlogs) createDirAll(path string, err error) {
	if os.IsNotExist(err) {
		_ = os.MkdirAll(path, 0777)
	}
}
