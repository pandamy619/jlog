package logging

import (
	"fmt"
	"os"
	"testing"
)


func openFile() (*os.File, error){
	file, err := os.OpenFile("./tmp/test.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644)
	return file, err
}

func TestNewLog(t *testing.T) {
	fmt.Println("Start test TestNewLog")

	file, err := openFile()
	defer file.Close()
	if err != nil {
		t.Error("file opening error")
		os.Exit(1)
	}
	NewLog(file)

	fmt.Println("Stop test TestNewLog")
}

func TestLogger_Output(t *testing.T) {
	fmt.Println("Start TestLogger_Output")

	file, err := openFile()
	defer file.Close()
	if err != nil {
		t.Error("file opening error")
		os.Exit(1)
	}
	l := NewLog(file)

	err = l.Output("test row")
	if err != nil {
		t.Error("error write to file")
	}

	fmt.Println("Stop TestLogger_Output")
}