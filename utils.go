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

//createDirs creates a directory named path
func createDirs(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(path, 0777)
	}
}

func writeToJson(name string, makeData interface{}) {
	file, _ := os.OpenFile(name, os.O_CREATE, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	_ = encoder.Encode(makeData)
}

func (l *Jlogs) outFile() {
	path := fmt.Sprintf("%s/%s", l.location, l.Pkg)
	createDirs(path)
	// TODO временное решение
	filename := time.Now().Format("2006-01-02T15:04:05")
	_ = ioutil.WriteFile(
		fmt.Sprintf("%s/%s/%s.json", l.location, l.Pkg, filename),
		l.Report(),
		0777,
	)
}
