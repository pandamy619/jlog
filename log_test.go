package logging

import (
	"fmt"
	"os"
	"testing"
	"time"
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
	NewLog(time.Now().Format("2006-01-02T15:04:05"), file)

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
	l := NewLog(time.Now().Format("2006-01-02T15:04:05"), file)

	err = l.Output("test row 1")
	if err != nil {
		t.Error("error write to file")
	}
	err = l.Output("test row 2")
	fmt.Println("Stop TestLogger_Output")
}

func TestLogger_Print(t *testing.T) {
	fmt.Println("Start test TestLogger_Info")

	file, err := openFile()
	defer file.Close()
	if err != nil {
		t.Error("file opening error")
		os.Exit(1)
	}

	l := NewLog(time.Now().Format("2006-01-02T15:04:05"), file)
	l.WithFields(Fields{
		"package": "logging",
		"function": "TestLogger_Info",
	})
	fmt.Println("")
	l.Info("message info")
	fmt.Println("")
	l.Warning("message warning")
	fmt.Println("")
	l.Error("message error")
	fmt.Println("")
	fmt.Println("Stop test TestLogger_Info")
}