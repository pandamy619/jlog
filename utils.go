package jlog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func structJson(obj interface{}) []byte{
	o, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return o
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
	_ = ioutil.WriteFile(filename, makeData,0777)
}
