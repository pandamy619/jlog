package jlog

import (
	"fmt"
	"os"
	"time"
)

// TODO оставить нормальные комментарии
func (l *Jlogs) makePath() string {
	return fmt.Sprintf("%s/%s", l.location, l.Pkg)
}

func (l *Jlogs) makeFilename() string {
	return fmt.Sprintf("%s.%s", time.Now().Format(timeFormat), l.consoleType)
}

// Creates a directory named path
func createDirs(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(path, 0777)
	}
}

// outFile write or append data to a file.
// If directory or file not exist, will be created.
func outFile(path string, filename string, data []byte) {
	createDirs(path)
	err := writeFile(path+"/"+filename, data, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

// writeFile writes data to a file named by filename.
// If the file does not exist, the file will be created.
// If the file exist, data will be appended to file
func writeFile(filename string, data []byte, perm os.FileMode) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)

	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err1 := file.Close(); err == nil {
		err = err1
	}
	return err
}
